/*
Copyright 2019-2024 Google LLC

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
	"cloud.google.com/go/asset/apiv1/assetpb"
	container "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	"cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	cloudresourcemanager2 "google.golang.org/api/cloudresourcemanager/v2"
	cloudresourcemanager3 "google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/logging/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/servicemanagement/v1"
)

const (
	LifecycleStateActiveRequested = "ACTIVE"
	TargetExcludedLabels          = "TARGET_EXCLUDED_LABELS"
	TargetIncludedLabels          = "TARGET_INCLUDED_LABELS"
	CleanUpTagKeys                = "CLEAN_UP_TAG_KEYS"
	CleanUpSCCNotfi               = "CLEAN_UP_SCC_NOTIFICATIONS"
	TargetExcludedTagKeys         = "TARGET_EXCLUDED_TAGKEYS"
	TargetIncludedSCCNotfis       = "TARGET_INCLUDED_SCC_NOTIFICATIONS"
	TargetFolderId                = "TARGET_FOLDER_ID"
	TargetOrganizationId          = "TARGET_ORGANIZATION_ID"
	MaxProjectAgeHours            = "MAX_PROJECT_AGE_HOURS"
	targetFolderRegexp            = `^[0-9]+$`
	targetOrganizationRegexp      = `^[0-9]+$`
	billingAccountRegex           = `^[0-9A-Z][-0-9A-Z]{18}[0-9A-Z]$`
	SCCNotificationsPageSize      = "SCC_NOTIFICATIONS_PAGE_SIZE"
	CleanUpCaiFeeds               = "CLEAN_UP_CAI_FEEDS"
	TargetIncludedFeeds           = "TARGET_INCLUDED_FEEDS"
	BillingAccount                = "BILLING_ACCOUNT"
	CleanUpBillingSinks           = "CLEAN_UP_BILLING_SINKS"
	TargetBillingSinks            = "TARGET_BILLING_SINKS"
	BillingSinksPageSize          = "BILLING_SINKS_PAGE_SIZE"
)

var (
	logger                 = log.New(os.Stdout, "", 0)
	excludedLabelsMap      = getLabelsMapFromEnv(TargetExcludedLabels)
	includedLabelsMap      = getLabelsMapFromEnv(TargetIncludedLabels)
	cleanUpTagKeys         = getBoolFromEnv(CleanUpTagKeys)
	cleanUpSCCNotfi        = getBoolFromEnv(CleanUpSCCNotfi)
	excludedTagKeysList    = getTagKeysListFromEnv(TargetExcludedTagKeys)
	includedSCCNotfisList  = getRegexListFromEnv(TargetIncludedSCCNotfis)
	resourceCreationCutoff = getOldTime(getIntFromEnv(MaxProjectAgeHours) * 60 * 60)
	rootFolderId           = getCorrectFolderIdOrTerminateExecution()
	organizationId         = getCorrectOrganizationIdOrTerminateExecution()
	sccPageSize            = int32(getIntFromEnv(SCCNotificationsPageSize))
	cleanUpCaiFeeds        = getBoolFromEnv(CleanUpCaiFeeds)
	includedFeedsList      = getRegexListFromEnv(TargetIncludedFeeds)
	billingAccount         = getBillingAccountOrTerminateExecution()
	cleanUpBillingSinks    = getBoolFromEnv(CleanUpBillingSinks)
	billingSinksPageSize   = getIntFromEnv(BillingSinksPageSize)
	targetBillingSinks     = getRegexListFromEnv(TargetBillingSinks)
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

func checkIfNameIncluded(name string, reg []*regexp.Regexp) bool {
	if len(reg) == 0 {
		return false
	}
	for _, regex := range reg {
		if regex.MatchString(name) {
			return true
		}
	}
	return false
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

func getRegexListFromEnv(envVariableName string) []*regexp.Regexp {
	var compiledRegEx []*regexp.Regexp
	envListVar := os.Getenv(envVariableName)
	logger.Printf("Try to get [%s] list", envVariableName)
	if envListVar == "" {
		logger.Printf("No value for [%s] list provided.", envVariableName)
		return compiledRegEx
	}

	var regexList []string
	err := json.Unmarshal([]byte(envListVar), &regexList)
	if err != nil {
		logger.Printf("Failed to get Regex list from [%s] env variable, error [%s]", envVariableName, err.Error())
		return compiledRegEx
	} else {
		logger.Printf("Got Regex list [%s] from [%s] env variable", regexList, envVariableName)
	}

	// build Regexes
	for _, r := range regexList {
		result, err := regexp.Compile(r)
		if err != nil {
			logger.Printf("Invalid regular expression [%s] for [%s]", r, envVariableName)
		} else {
			compiledRegEx = append(compiledRegEx, result)
		}
	}
	return compiledRegEx
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

func getBoolFromEnv(envVariableName string) bool {
	envVariableNameVal, exists := os.LookupEnv(envVariableName)
	if !exists {
		logger.Fatalf("Environment variable [%s] not set, set the environment variable and try again.", envVariableName)
	}
	result, err := strconv.ParseBool(envVariableNameVal)
	if err != nil {
		logger.Fatalf("Invalid bool value [%s], specify correct value for environment variable [%s] and try again.", envVariableNameVal, envVariableName)
	}
	return result
}

func getIntFromEnv(envVariableName string) int64 {
	envVariableStr := os.Getenv(envVariableName)
	intValue, err := strconv.ParseInt(envVariableStr, 10, 0)
	if err != nil {
		logger.Fatalf("Could not convert [%s] to integer. Specify correct value for environment variable [%s] and try again.", envVariableStr, envVariableName)
	}
	return intValue
}

func getCorrectFolderIdOrTerminateExecution() string {
	targetFolderIdString := os.Getenv(TargetFolderId)
	matched, err := regexp.MatchString(targetFolderRegexp, targetFolderIdString)
	if err != nil || !matched {
		logger.Fatalf("Invalid folder id [%s], specify correct value and try again.", targetFolderIdString)
	}
	return targetFolderIdString
}

func getBillingAccountOrTerminateExecution() string {
	billingAccountVal := os.Getenv(BillingAccount)
	if billingAccountVal == "" {
		if cleanUpBillingSinks {
			logger.Fatal("If billing account sink clean up is enabled, billing account id should not be empty, specify correct value and try again.")
		}
		return billingAccountVal
	}
	matched, err := regexp.MatchString(billingAccountRegex, billingAccountVal)
	if err != nil || !matched {
		logger.Fatalf("Invalid billing account id [%s], specify correct value and try again.", billingAccountVal)
	}
	return billingAccountVal
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

func getBillingAccountSinkServiceOrTerminateExecution(ctx context.Context, client *http.Client) *logging.BillingAccountsSinksService {
	loggingService, err := logging.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Fatalf("Failed to get Logging Sink Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Logging Sink Service")
	return loggingService.BillingAccounts.Sinks
}

func getSCCNotificationServiceOrTerminateExecution(ctx context.Context) *securitycenter.Client {
	logger.Println("Try to get SCC Notification Service")
	securitycenterClient, err := securitycenter.NewClient(ctx)
	if err != nil {
		logger.Fatalf("Failed to get SCC Notification Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got SCC Notification Service")
	return securitycenterClient
}

func getAssetServiceOrTerminateExecution(ctx context.Context) *asset.Client {
	logger.Println("Try to get Asset Service")
	assetService, err := asset.NewClient(ctx)
	if err != nil {
		logger.Fatalf("Failed to get Asset Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Asset Service")
	return assetService
}

func getContainerServiceOrTerminateExecution(ctx context.Context) *container.ClusterManagerClient {
	logger.Println("Try to get Container Service")
	containerService, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		logger.Fatalf("Failed to get Container Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Container Service")
	return containerService
}

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
	sccService := getSCCNotificationServiceOrTerminateExecution(ctx)
	tagValuesService := getTagValuesServiceOrTerminateExecution(ctx, client)
	feedsService := getAssetServiceOrTerminateExecution(ctx)
	billingSinkService := getBillingAccountSinkServiceOrTerminateExecution(ctx, client)
	firewallPoliciesService := getFirewallPoliciesServiceOrTerminateExecution(ctx, client)
	endpointService := getServiceManagementServiceOrTerminateExecution(ctx, client)
	containerService := getContainerServiceOrTerminateExecution(ctx)

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

	billingSinkAgeFilter := func(logSink *logging.LogSink) bool {
		createdAt, err := time.Parse(time.RFC3339, logSink.CreateTime)
		if err != nil {
			logger.Printf("Failed to parse CreateTime for tagKey [%s], skipping it, error [%s]", logSink.ResourceName, err.Error())
			return false
		}
		return createdAt.Before(resourceCreationCutoff)
	}

	projectDeleteRequestedFilter := func(projectID string) bool {
		p, err := cloudResourceManagerService.Projects.Get(projectID).Context(ctx).Do()
		if err != nil {
			logger.Printf("Failed to get project [%s], error [%s]", projectID, err.Error())
			return false
		}
		if p.LifecycleState == "DELETE_REQUESTED" {
			return true
		}
		return false
	}

	removeSCCNotifications := func(organization string) {
		logger.Printf("Try to remove SCC Notifications from organization [%s]", organization)
		req := &securitycenterpb.ListNotificationConfigsRequest{
			Parent:   fmt.Sprintf("organizations/%s", organization),
			PageSize: sccPageSize,
		}
		it := sccService.ListNotificationConfigs(ctx, req)
		for {
			resp, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				logger.Printf("failed to list SCC notifications, error [%s]", err.Error())
				break
			}
			projectID := strings.Split(resp.PubsubTopic, "/")[1]
			if checkIfNameIncluded(resp.Name, includedSCCNotfisList) && projectDeleteRequestedFilter(projectID) {
				delReq := &securitycenterpb.DeleteNotificationConfigRequest{
					Name: resp.Name,
				}
				err = sccService.DeleteNotificationConfig(ctx, delReq)
				if err != nil {
					logger.Printf("failed to delete SCC notification [%s], error [%s]", resp.Name, err.Error())
				} else {
					logger.Printf("SCC notification [%s] deleted", resp.Name)
				}
			}
		}
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

	removeFeedsByName := func(organization string) {
		logger.Printf("Try to remove feeds from organization [%s]", organization)

		req := &assetpb.ListFeedsRequest{
			Parent: fmt.Sprintf("organizations/%s", organization),
		}

		resp, err := feedsService.ListFeeds(ctx, req)
		if err != nil {
			logger.Printf("Failed to list Feeds, error [%s]", err.Error())
			return
		}

		for _, feed := range resp.Feeds {
			projectID := strings.Split(feed.FeedOutputConfig.GetPubsubDestination().Topic, "/")[1]
			if checkIfNameIncluded(feed.Name, includedFeedsList) && projectDeleteRequestedFilter(projectID) {
				delReq := &assetpb.DeleteFeedRequest{
					Name: feed.Name,
				}
				err := feedsService.DeleteFeed(ctx, delReq)
				if err != nil {
					logger.Printf("Failed to remove the feed [%s], error [%s]", feed.Name, err.Error())
				} else {
					logger.Printf("Feed [%s] successfully removed.", feed.Name)
				}
			}
		}
	}

	removeBillingSinks := func(billing string) {
		logger.Printf("Try to remove billing account log sinks from billing account [%s]", billing)
		parent := fmt.Sprintf("billingAccounts/%s", billing)
		sinkList, err := billingSinkService.List(parent).PageSize(billingSinksPageSize).Context(ctx).Do()
		if err != nil {
			logger.Printf("Failed to list billing account log sinks from billing account [%s], error [%s]", billing, err.Error())
			return
		}
		for _, sink := range sinkList.Sinks {
			if sink.Name != "_Required" && sink.Name != "_Default" && billingSinkAgeFilter(sink) && checkIfNameIncluded(sink.ResourceName, targetBillingSinks) {
				_, err = billingSinkService.Delete(sink.ResourceName).Context(ctx).Do()
				if err != nil {
					logger.Printf("Failed to delete billing account log sink [%s] from billing account [%s], error [%s]", sink.ResourceName, billing, err.Error())
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

	removeProjectClusters := func(projectId string) int {
		logger.Printf("Try to remove clusters for [%s]", projectId)
		reqLCR := &containerpb.ListClustersRequest{Parent: fmt.Sprintf("projects/%s/locations/*", projectId)}
		listResponse, err := containerService.ListClusters(ctx, reqLCR)
		if err != nil {
			logger.Printf("Failed to list clusters for [%s], error [%s]", projectId, err.Error())
			return 0
		}

		logger.Printf("Got [%d] clusters for project [%s]", len(listResponse.Clusters), projectId)
		if len(listResponse.Clusters) == 0 {
			return 0
		}

		var pendingDeletion int = 0
		for _, cluster := range listResponse.Clusters {
			switch clusterStatus := cluster.Status.String(); clusterStatus {
			case "DEGRADED":
				fallthrough
			case "RUNNING":
				logger.Printf("Deleting cluster %s status: %s", cluster.Name, clusterStatus)
				reqDCR := &containerpb.DeleteClusterRequest{Name: fmt.Sprintf("projects/%s/locations/%s/clusters/%s", projectId, cluster.Location, cluster.Name)}
				_, err := containerService.DeleteCluster(ctx, reqDCR)
				if err != nil {
					logger.Printf("Failed to delete cluster [%s] for [%s], error [%s]", cluster.Name, projectId, err.Error())
				} else {
					pendingDeletion++
				}
			case "PROVISIONING":
				fallthrough
			case "RECONCILING":
				fallthrough
			case "STOPPING":
				logger.Printf("Deferring cluster %s status: %s", cluster.Name, clusterStatus)
				pendingDeletion++
			default:
				logger.Printf("Ignoring cluster %s status: %s", cluster.Name, clusterStatus)
			}
		}
		return pendingDeletion
	}

	removeProjectEndpoints := func(projectId string) {
		logger.Printf("Try to remove endpoints for [%s]", projectId)
		listResponse, err := endpointService.Services.List().ProducerProjectId(projectId).Do()
		if err != nil {
			logger.Printf("Failed to list services for [%s], error [%s]", projectId, err.Error())
			return
		}

		logger.Printf("Got [%d] services for the project [%s]", len(listResponse.Services), projectId)
		if len(listResponse.Services) == 0 {
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
		if clusters := removeProjectClusters(projectId); clusters != 0 {
			logger.Printf("Defer removing project [%s], %d clusters marked for deletion", projectId, clusters)
		}
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

	// only delete Security Command Center notifications from deleted projects
	if cleanUpSCCNotfi {
		removeSCCNotifications(organizationId)
	}

	// Only delete Feeds from deleted projects
	if cleanUpCaiFeeds {
		removeFeedsByName(organizationId)
	}

	if cleanUpBillingSinks {
		removeBillingSinks(billingAccount)
	}
}

func CleanUpProjects(ctx context.Context, m PubSubMessage) error {
	invoke(ctx)
	return nil
}
