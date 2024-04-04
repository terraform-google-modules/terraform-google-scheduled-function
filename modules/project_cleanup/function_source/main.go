/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package project_cleanup

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	asset "cloud.google.com/go/asset/apiv1"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	cloudresourcemanager2 "google.golang.org/api/cloudresourcemanager/v2"
	cloudresourcemanager3 "google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/servicemanagement/v1"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	LifecycleStateActiveRequested = "ACTIVE"
	TargetExcludedLabels          = "TARGET_EXCLUDED_LABELS"
	TargetIncludedLabels          = "TARGET_INCLUDED_LABELS"
	CleanUpTagKeys                = "CLEAN_UP_TAG_KEYS"
	TargetExcludedTagKeys         = "TARGET_EXCLUDED_TAGKEYS"
	TargetFolderId                = "TARGET_FOLDER_ID"
	TargetOrganizationId          = "TARGET_ORGANIZATION_ID"
	MaxProjectAgeHours            = "MAX_PROJECT_AGE_HOURS"
	targetFolderRegexp            = `^[0-9]+$`
	targetOrganizationRegexp      = `^[0-9]+$`
	CleanUpCaiFeeds               = "CLEAN_UP_CAI_FEEDS"    ///r
	TargetIncludedFeeds           = "TARGET_INCLUDED_FEEDS" ///r
)

var (
	logger                 = log.New(os.Stdout, "", 0)
	excludedLabelsMap      = getLabelsMapFromEnv(TargetExcludedLabels)
	includedLabelsMap      = getLabelsMapFromEnv(TargetIncludedLabels)
	cleanUpTagKeys         = getCleanUpTagKeysOrTerminateExecution()
	excludedTagKeysList    = getTagKeysListFromEnv(TargetExcludedTagKeys)
	resourceCreationCutoff = getOldTime(int64(getCorrectMaxAgeInHoursOrTerminateExecution()) * 60 * 60)
	rootFolderId           = getCorrectFolderIdOrTerminateExecution()
	organizationId         = getCorrectOrganizationIdOrTerminateExecution()
	cleanUpCaiFeeds        = getCleanUpFeedsOrTerminateExecution()    ///r
	includedFeedsList      = getFeedsListFromEnv(TargetIncludedFeeds) ///r

)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type FolderRecursion func(*cloudresourcemanager2.Folder, FolderRecursion)

func activeProjectFilter(project *cloudresourcemanager.Project) bool {
	return project.LifecycleState == LifecycleStateActiveRequested
}

func getOldTime(i int64) time.Time {
	return time.Unix(time.Now().Unix()-i, 0)
}

// isRetryableError checks if an error can be retried based on err code
func isRetryableError(e error) bool {
	gerr, ok := e.(*googleapi.Error)
	if !ok {
		return false
	}
	if gerr.Code == 429 || gerr.Code == 500 || gerr.Code == 502 || gerr.Code == 503 {
		logger.Printf("Got retryable err %d: %s", gerr.Code, gerr.Message)
		return true
	}
	logger.Printf("Got non retryable err %d: %s", gerr.Code, gerr.Message)
	return false
}

func retry(retryFunc func() error, tries int, duration time.Duration) error {
	err := retryFunc()
	if err == nil {
		return nil
	}
	if tries < 1 {
		return fmt.Errorf("Exhausted retries: %v", err)
	}
	if isRetryableError(err) {
		time.Sleep(duration)
		tries--
		// retry with exponential backoff
		return retry(retryFunc, tries, 2*duration)
	}
	return err
}

func processProjectsResponsePage(removeProjectById func(projectId string)) func(page *cloudresourcemanager.ListProjectsResponse) error {
	excludeProjectByOneOfLabelsFilter := func(project *cloudresourcemanager.Project) bool {
		return !checkIfAtLeastOneLabelPresentIfAny(project, excludedLabelsMap, true)
	}

	includeProjectByOneOfLabelsFilter := func(project *cloudresourcemanager.Project) bool {
		return checkIfAtLeastOneLabelPresentIfAny(project, includedLabelsMap, false)
	}

	ageFilter := func(project *cloudresourcemanager.Project) bool {
		projectCreatedAt, err := time.Parse(time.RFC3339, project.CreateTime)
		if err != nil {
			logger.Printf("Failed to parse CreateTime for [%s], skipping it, error [%s]", project.Name, err.Error())
			return false
		}
		return projectCreatedAt.Before(resourceCreationCutoff)
	}

	combinedProjectFilter := func(project *cloudresourcemanager.Project) bool {
		return activeProjectFilter(project) && ageFilter(project) && includeProjectByOneOfLabelsFilter(project) && excludeProjectByOneOfLabelsFilter(project)
	}

	return func(page *cloudresourcemanager.ListProjectsResponse) error {
		for _, project := range page.Projects {
			if combinedProjectFilter(project) {
				projectId := project.ProjectId
				removeProjectById(projectId)
			}
		}
		return nil
	}
}

func getCorrectMaxAgeInHoursOrTerminateExecution() int64 {
	maxAgeInHoursStr := os.Getenv(MaxProjectAgeHours)
	maxAgeInHours, err := strconv.ParseInt(os.Getenv(MaxProjectAgeHours), 10, 0)
	if err != nil {
		logger.Fatalf("Could not convert [%s] to integer. Specify correct value for environment variable [%s] and try again.", maxAgeInHoursStr, MaxProjectAgeHours)
	}
	return maxAgeInHours
}

func checkIfAtLeastOneLabelPresentIfAny(project *cloudresourcemanager.Project, labels map[string]string, isExcludeCheck bool) bool {
	if len(labels) == 0 {
		return !isExcludeCheck
	}
	result := false
	projectLabels := project.Labels
	for key, value := range labels {
		if !result {
			result = projectLabels[key] == value
		}
	}
	return result
}

func checkIfTagKeyShortNameExcluded(shortName string, excludedTagKeys []string) bool {
	if len(excludedTagKeys) == 0 {
		return false
	}
	for _, name := range excludedTagKeys {
		if shortName == name {
			return true
		}
	}
	return false
}

///r
func checkIfCaiFeedsShortNameIncluded(shortName string, IncludedFeeds []string) bool {
	if len(IncludedFeeds) > 0 {
		return true
	}
	for _, name := range IncludedFeeds {
		if shortName == name {
			return true
		}
	}
	return false
}

///r

func getLabelsMapFromEnv(envVariableName string) map[string]string {
	targetExcludedLabels := os.Getenv(envVariableName)
	logger.Println("Try to get labels map")
	labels := make(map[string]string)

	if targetExcludedLabels == "" {
		logger.Printf("No labels provided.")
		return nil
	}

	err := json.Unmarshal([]byte(targetExcludedLabels), &labels)
	if err != nil {
		logger.Printf("Failed to get labels map from [%s] env variable, error [%s]", envVariableName, err.Error())
	} else {
		logger.Printf("Got labels map [%s] from [%s] env variable", labels, envVariableName)
	}
	return labels
}

func getTagKeysListFromEnv(envVariableName string) []string {
	targetExcludedTagKeys := os.Getenv(envVariableName)
	logger.Println("Try to get Tag Keys list")
	if targetExcludedTagKeys == "" {
		logger.Printf("No Tag Keys provided.")
		return nil
	}

	var tagKeys []string
	err := json.Unmarshal([]byte(targetExcludedTagKeys), &tagKeys)
	if err != nil {
		logger.Printf("Failed to get Tag Keys list from [%s] env variable, error [%s]", envVariableName, err.Error())
	} else {
		logger.Printf("Got Tag Keys list [%s] from [%s] env variable", tagKeys, envVariableName)
	}
	return tagKeys
}

func getCleanUpTagKeysOrTerminateExecution() bool {
	cleanUpTagKeys, exists := os.LookupEnv(CleanUpTagKeys)
	if !exists {
		logger.Fatalf("Clean up Tag Keys environment variable [%s] not set, set the environment variable and try again.", CleanUpTagKeys)
	}
	result, err := strconv.ParseBool(cleanUpTagKeys)
	if err != nil {
		logger.Fatalf("Invalid Clean up Tag Keys value [%s], specify correct value for environment variable [%s] and try again.", cleanUpTagKeys, CleanUpTagKeys)
	}
	return result
}

///r
func getFeedsListFromEnv(envVariableName string) []string {
	TargetIncludedFeeds := os.Getenv(envVariableName)
	logger.Println("Try to get Cai Feeds list")
	if TargetIncludedFeeds == "" {
		logger.Printf("No Cai Feeds provided.")
		return nil
	}

	var caiFeeds []string
	err := json.Unmarshal([]byte(TargetIncludedFeeds), &caiFeeds)
	if err != nil {
		logger.Printf("Failed to get Cai Feeds list from [%s] env variable, error [%s]", envVariableName, err.Error())
	} else {
		logger.Printf("Got Cai Feeds list [%s] from [%s] env variable", caiFeeds, envVariableName)
	}
	return caiFeeds
}

func getCleanUpFeedsOrTerminateExecution() bool {
	cleanUpCaiFeeds, exists := os.LookupEnv(CleanUpCaiFeeds)
	if !exists {
		logger.Fatalf("Clean up CaiFeeds environment variable [%s] not set, set the environment variable and try again.", CleanUpCaiFeeds)
	}
	result, err := strconv.ParseBool(cleanUpCaiFeeds)
	if err != nil {
		logger.Fatalf("Invalid Clean up CaiFeeds value [%s], specify correct value for environment variable [%s] and try again.", cleanUpCaiFeeds, CleanUpCaiFeeds)
	}
	return result
}

///r

func getCorrectFolderIdOrTerminateExecution() string {
	targetFolderIdString := os.Getenv(TargetFolderId)
	matched, err := regexp.MatchString(targetFolderRegexp, targetFolderIdString)
	if err != nil || !matched {
		logger.Fatalf("Invalid folder id [%s], specify correct value and try again.", targetFolderIdString)
	}
	return targetFolderIdString
}

func getCorrectOrganizationIdOrTerminateExecution() string {
	targetOrganizationIdString := os.Getenv(TargetOrganizationId)
	matched, err := regexp.MatchString(targetOrganizationRegexp, targetOrganizationIdString)
	if err != nil || !matched {
		logger.Fatalf("Invalid organization id [%s], specify correct value and try again.", targetOrganizationIdString)
	}
	return targetOrganizationIdString
}

func getServiceManagementServiceOrTerminateExecution(ctx context.Context, client *http.Client) *servicemanagement.APIService {
	service, err := servicemanagement.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Fatalf("Failed to get service management API client with error [%s], terminate execution", err.Error())
	}
	return service
}

func getResourceManagerServiceOrTerminateExecution(ctx context.Context, client *http.Client) *cloudresourcemanager.Service {
	logger.Println("Try to get Cloud Resource Manager")
	cloudResourceManagerService, err := cloudresourcemanager.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Fatalf("Failed to get Cloud Resource Manager with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Cloud Resource Manager")
	return cloudResourceManagerService
}

func getFolderServiceOrTerminateExecution(ctx context.Context, client *http.Client) *cloudresourcemanager2.FoldersService {
	logger.Println("Try to get Folders Service")
	cloudResourceManagerService, err := cloudresourcemanager2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Fatalf("Failed to get Folders Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Folders Service")
	return cloudResourceManagerService.Folders
}

func getTagKeysServiceOrTerminateExecution(ctx context.Context, client *http.Client) *cloudresourcemanager3.TagKeysService {
	logger.Println("Try to get TagKeys Service")
	cloudResourceManagerService, err := cloudresourcemanager3.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Fatalf("Failed to get TagKeys Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got TagKeys Service")
	return cloudResourceManagerService.TagKeys
}

func getTagValuesServiceOrTerminateExecution(ctx context.Context, client *http.Client) *cloudresourcemanager3.TagValuesService {
	logger.Println("Try to get TagValues Service")
	cloudResourceManagerService, err := cloudresourcemanager3.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Fatalf("Failed to get TagValues Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got TagValues Service")
	return cloudResourceManagerService.TagValues
}

///r
//https://pkg.go.dev/cloud.google.com/go/asset/apiv1#Client.ListAssets
func getListFeedsRequestOrTerminateExecution(ctx context.Context, client *http.Client, organizationId string) *assetpb.ListFeedsRequest {
	///func getListFeedsRequestOrTerminateExecution(ctx context.Context, client *http.Client, organizationId string) (*assetpb.ListFeedsRequest, error) {

	logger.Println("Creating ListFeedsRequest")

	assetClient, err := asset.NewClient(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Fatalf("Failed to create Asset client with error [%s], terminate execution", err.Error())
	}

	resp, err := assetClient.ListFeeds(ctx, &assetpb.ListFeedsRequest{
		Parent: fmt.Sprintf("organizations/%s", organizationId),
	})
	if err != nil {
		logger.Fatalf("Failed to list feeds with error [%s], terminate execution", err.Error())
	}

	for _, feed := range resp.Feeds {
		fmt.Println(feed)
	}

	req := &assetpb.ListFeedsRequest{
		Parent: fmt.Sprintf("organizations/%s", organizationId),
	}

	return req
	///return req, nil
}

///r

func getFirewallPoliciesServiceOrTerminateExecution(ctx context.Context, client *http.Client) *compute.FirewallPoliciesService {
	logger.Println("Try to get Firewall Policies Service")
	computeService, err := compute.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Fatalf("Failed to get Firewall Policies Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Firewall Policies Service")
	return computeService.FirewallPolicies
}

func initializeGoogleClient(ctx context.Context) *http.Client {
	logger.Println("Try to initialize Google client")
	client, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		logger.Fatalf("Failed to initialize Google client with error [%s], terminate execution", err.Error())
	}
	logger.Println("Initialized Google client")
	return client
}

func invoke(ctx context.Context) {
	client := initializeGoogleClient(ctx)
	cloudResourceManagerService := getResourceManagerServiceOrTerminateExecution(ctx, client)
	folderService := getFolderServiceOrTerminateExecution(ctx, client)
	tagKeyService := getTagKeysServiceOrTerminateExecution(ctx, client)
	tagValuesService := getTagValuesServiceOrTerminateExecution(ctx, client)
	listFeedsRequest := getListFeedsRequestOrTerminateExecution(ctx, client, organizationId)
	//listFeedsRequest, err := getListFeedsRequestOrTerminateExecution(ctx, client, organizationId)
	// if err != nil {
	// 	logger.Fatalf("Failed to get ListFeedsRequest with error [%s], terminate execution", err.Error())
	// }
	firewallPoliciesService := getFirewallPoliciesServiceOrTerminateExecution(ctx, client)
	endpointService := getServiceManagementServiceOrTerminateExecution(ctx, client)

	removeLien := func(name string) {
		logger.Printf("Try to remove lien [%s]", name)
		_, err := cloudResourceManagerService.Liens.Delete(name).Context(ctx).Do()
		if err != nil {
			logger.Printf("Failed to remove lien [%s], error [%s]", name, err.Error())
		} else {
			logger.Printf("Removed lien [%s]", name)
		}
	}

	tagKeyAgeFilter := func(tagKey *cloudresourcemanager3.TagKey) bool {
		tagKeyCreatedAt, err := time.Parse(time.RFC3339, tagKey.CreateTime)
		if err != nil {
			logger.Printf("Failed to parse CreateTime for tagKey [%s], skipping it, error [%s]", tagKey.Name, err.Error())
			return false
		}
		return tagKeyCreatedAt.Before(resourceCreationCutoff)
	}

	removeTagValues := func(tagKey string) {
		logger.Printf("Try to remove Tag Values from TagKey [%s]", tagKey)
		tagValuesList, err := tagValuesService.List().Parent(tagKey).Context(ctx).Do()
		if err != nil {
			logger.Printf("Failed to list Tag values from TagKey [%s], error [%s]", tagKey, err.Error())
			return
		}
		for _, tagValue := range tagValuesList.TagValues {
			_, err := tagValuesService.Delete(tagValue.Name).Context(ctx).Do()
			if err != nil {
				logger.Printf("Failed to delete tagValue from TagKey [%s], error [%s]", tagKey, err.Error())
			}
		}
	}

	removeTagKeys := func(organization string) {
		logger.Printf("Try to remove Tag Keys from organization [%s]", organization)
		parent := fmt.Sprintf("organizations/%s", organization)
		tagKeysList, err := tagKeyService.List().Parent(parent).Context(ctx).Do()
		if err != nil {
			logger.Printf("Failed to list Tag Keys from organization [%s], error [%s]", organization, err.Error())
			return
		}
		for _, tagKey := range tagKeysList.TagKeys {
			if !checkIfTagKeyShortNameExcluded(tagKey.ShortName, excludedTagKeysList) && tagKeyAgeFilter(tagKey) {
				removeTagValues(tagKey.Name)
				_, err := tagKeyService.Delete(tagKey.Name).Context(ctx).Do()
				if err != nil {
					logger.Printf("Failed to delete tagKey from organization [%s], error [%s]", organization, err.Error())
				}
			}
		}
	}

	removeFirewallPolicies := func(folder string) {
		logger.Printf("Try to remove Firewall Policies from folder [%s]", folder)
		firewallPolicyList, err := firewallPoliciesService.List().ParentId(folder).Context(ctx).Do()
		if err != nil {
			logger.Printf("Failed to list Firewall Policies from folder [%s], error [%s]", folder, err.Error())
			return
		}
		for _, policy := range firewallPolicyList.Items {
			for _, association := range policy.Associations {
				_, err := firewallPoliciesService.RemoveAssociation(policy.Name).Name(association.Name).Context(ctx).Do()
				if err != nil {
					logger.Printf("Failed to Remove Association for Firewall Policies from folder [%s], error [%s]", folder, err.Error())
				}
			}
			_, err := firewallPoliciesService.Delete(policy.Name).Context(ctx).Do()
			if err != nil {
				logger.Printf("Failed to delete Firewall Policy [%s] from folder [%s], error [%s]", policy.Name, folder, err.Error())
			}
		}
	}

	removeProjectById := func(projectId string) error {
		_, err := cloudResourceManagerService.Projects.Delete(projectId).Context(ctx).Do()
		return err
	}

	removeProjectEndpoints := func(projectId string) {
		logger.Printf("Try to remove endpoints for [%s]", projectId)
		listResponse, err := endpointService.Services.List().ProducerProjectId(projectId).Do()
		if err != nil {
			logger.Printf("Failed to list services for [%s], error [%s]", projectId, err.Error())
			return
		}

		if len(listResponse.Services) <= 1 {
			return
		}

		for _, service := range listResponse.Services {
			logger.Printf("Try to remove service: %s", service.ServiceName)
			_, err = endpointService.Services.Delete(service.ServiceName).Do()
			if err != nil {
				logger.Printf("Failed to delete service [%s] for [%s], error [%s]", service.ServiceName, projectId, err.Error())
			}
		}

		// wait for services to complete deletion
		time.Sleep(10 * time.Second)
	}

	cleanupProjectById := func(projectId string) {
		logger.Printf("Try to remove project [%s]", projectId)
		err := removeProjectById(projectId)
		if err != nil {
			removeProjectEndpoints(projectId)
			err = removeProjectById(projectId)
		}
		if err != nil {
			logger.Printf("Failed to remove project [%s], error [%s]", projectId, err.Error())
		} else {
			logger.Printf("Removed project [%s]", projectId)
		}
	}

	removeProjectWithLiens := func(projectId string) {
		logger.Printf("Try to get all liens for the project [%s]", projectId)
		parent := fmt.Sprintf("projects/%s", projectId)
		req := cloudResourceManagerService.Liens.List().Parent(parent)
		if err := req.Pages(ctx, func(page *cloudresourcemanager.ListLiensResponse) error {
			logger.Printf("Got [%d] liens for the project [%s]", len(page.Liens), projectId)
			for _, lien := range page.Liens {
				removeLien(lien.Name)
			}
			cleanupProjectById(projectId)
			return nil
		}); err != nil {
			logger.Printf("Failed to get all liens for the project [%s], error [%s]", projectId, err.Error())
		}
	}

	removeProjectsInFolder := func(folderId string) {
		localFolderId := strings.Replace(folderId, "folders/", "", 1)
		logger.Printf("Try to get projects from folder with id [%s] and process them", localFolderId)
		requestFilter := fmt.Sprintf("parent.type:folder parent.id:%s", localFolderId)
		err := retry(func() (err error) {
			req := cloudResourceManagerService.Projects.List().Filter(requestFilter)
			err = req.Pages(ctx, processProjectsResponsePage(removeProjectWithLiens))
			return
		}, 5, time.Minute)
		if err != nil {
			logger.Printf("Failed to get projects for the folder with id [%s], error [%s]", localFolderId, err.Error())
		} else {
			logger.Printf("Got and processed all projects for the folder with id [%s]", localFolderId)
		}
	}

	folderAgeFilter := func(folder *cloudresourcemanager2.Folder) bool {
		folderCreatedAt, err := time.Parse(time.RFC3339, folder.CreateTime)
		if err != nil {
			logger.Printf("Failed to parse CreateTime for folder [%s], skipping it, error [%s]", folder.Name, err.Error())
			return false
		}
		return folderCreatedAt.Before(resourceCreationCutoff)
	}

	removeFolder := func(folder *cloudresourcemanager2.Folder) {
		folderId := folder.Name
		removeFirewallPolicies(folderId)
		logger.Printf("Try to delete folder with id [%s]", folderId)
		_, err := folderService.Delete(folderId).Do()
		if err != nil {
			logger.Printf("Failed to delete folder [%s], error [%s]", folderId, err.Error())
		} else {
			logger.Printf("Deleted folder [%s]", folderId)
		}
	}

	getSubFoldersAndRemoveProjectsFoldersRecursively := func(folder *cloudresourcemanager2.Folder, recursion FolderRecursion) {
		folderId := folder.Name
		listFoldersRequest := folderService.List().Parent(folderId).ShowDeleted(false)
		if err := listFoldersRequest.Pages(ctx, func(foldersResponse *cloudresourcemanager2.ListFoldersResponse) error {
			for _, folder := range foldersResponse.Folders {
				recursion(folder, recursion)
			}
			removeProjectsInFolder(folderId)
			if folder.Parent != fmt.Sprintf("folders/%s", rootFolderId) && folder.Name != fmt.Sprintf("folders/%s", rootFolderId) && folderAgeFilter(folder) {
				removeFolder(folder)
			}
			return nil
		}); err != nil {
			logger.Fatalf("Failed to get subfolders for the folder with id [%s], error [%s]", folderId, err.Error())
		}
	}

	rootFolderId := fmt.Sprintf("folders/%s", rootFolderId)
	rootFolder, err := folderService.Get(rootFolderId).Do()
	if err != nil {
		logger.Printf("Failed to get parent folder [%s], error [%s]", rootFolderId, err.Error())
	} else {
		getSubFoldersAndRemoveProjectsFoldersRecursively(rootFolder, getSubFoldersAndRemoveProjectsFoldersRecursively)
	}
	// Only Tag Keys whose values are not in use can be deleted.
	if cleanUpTagKeys {
		removeTagKeys(organizationId)
	}
	// Only Feeds whose values are not in use can be deleted.
	if cleanUpCaiFeeds {
		removeFeedKeys(organizationId)
	}
}

func CleanUpProjects(ctx context.Context, m PubSubMessage) error {
	invoke(ctx)
	return nil
}
