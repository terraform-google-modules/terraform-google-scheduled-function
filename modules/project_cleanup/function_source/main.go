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

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	cloudresourcemanager2 "google.golang.org/api/cloudresourcemanager/v2"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/servicemanagement/v1"
)

const (
	LifecycleStateActiveRequested = "ACTIVE"
	TargetExcludedLabels          = "TARGET_EXCLUDED_LABELS"
	TargetIncludedLabels          = "TARGET_INCLUDED_LABELS"
	TargetFolderId                = "TARGET_FOLDER_ID"
	MaxProjectAgeHours            = "MAX_PROJECT_AGE_HOURS"
	targetFolderRegexp            = `^[0-9]+$`
)

var (
	logger                 = log.New(os.Stdout, "", 0)
	excludedLabelsMap      = getLabelsMapFromEnv(TargetExcludedLabels)
	includedLabelsMap      = getLabelsMapFromEnv(TargetIncludedLabels)
	resourceCreationCutoff = getOldTime(int64(getCorrectMaxAgeInHoursOrTerminateExecution()) * 60 * 60)
	rootFolderId           = getCorrectFolderIdOrTerminateExecution()
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
			logger.Printf("Fail to parse CreateTime for [%s], skip it. Error [%s]", project.Name, err.Error())
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
		logger.Fatalf("Could not convert [%s] to integer. Specify correct value, Please.", maxAgeInHoursStr)
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
		logger.Printf("Fail to get labels map from [%s] env variable, error [%s]", envVariableName, err.Error())
	} else {
		logger.Printf("Got labels map [%s] from [%s] env variable", labels, envVariableName)
	}
	return labels
}

func getCorrectFolderIdOrTerminateExecution() string {
	targetFolderIdString := os.Getenv(TargetFolderId)
	matched, err := regexp.MatchString(targetFolderRegexp, targetFolderIdString)
	if err != nil || !matched {
		logger.Fatalf("Invalid folder id [%s]. Specify correct value, Please.", targetFolderIdString)
	}
	return targetFolderIdString
}

func getServiceManagementServiceOrTerminateExecution(client *http.Client) *servicemanagement.APIService {
	service, err := servicemanagement.New(client)
	if err != nil {
		logger.Fatalf("Failed to get service management API client with error [%s], terminate execution", err.Error())
	}
	return service
}

func getResourceManagerServiceOrTerminateExecution(client *http.Client) *cloudresourcemanager.Service {
	logger.Println("Try to get Cloud Resource Manager")
	cloudResourceManagerService, err := cloudresourcemanager.New(client)
	if err != nil {
		logger.Fatalf("Fail to get Cloud Resource Manager with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Cloud Resource Manager")
	return cloudResourceManagerService
}

func getFolderServiceOrTerminateExecution(client *http.Client) *cloudresourcemanager2.FoldersService {
	logger.Println("Try to get Folders Service")
	cloudResourceManagerService, err := cloudresourcemanager2.New(client)
	if err != nil {
		logger.Fatalf("Fail to get Folders Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Folders Service")
	return cloudResourceManagerService.Folders
}

func getFirewallPoliciesServiceOrTerminateExecution(client *http.Client) *compute.FirewallPoliciesService {
	logger.Println("Try to get Firewall Policies Service")
	computeService, err := compute.New(client)
	if err != nil {
		logger.Fatalf("Fail to get Firewall Policies Service with error [%s], terminate execution", err.Error())
	}
	logger.Println("Got Firewall Policies Service")
	return computeService.FirewallPolicies
}

func initializeGoogleClient(ctx context.Context) *http.Client {
	logger.Println("Try to initialize Google client")
	client, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		logger.Fatalf("Fail to initialize Google client with error [%s], terminate execution", err.Error())
	}
	logger.Println("Initialized Google client")
	return client
}

func invoke(ctx context.Context) {
	client := initializeGoogleClient(ctx)
	cloudResourceManagerService := getResourceManagerServiceOrTerminateExecution(client)
	folderService := getFolderServiceOrTerminateExecution(client)
	firewallPoliciesService := getFirewallPoliciesServiceOrTerminateExecution(client)
	endpointService := getServiceManagementServiceOrTerminateExecution(client)

	removeLien := func(name string) {
		logger.Printf("Try to remove lien [%s]", name)
		_, err := cloudResourceManagerService.Liens.Delete(name).Context(ctx).Do()
		if err != nil {
			logger.Printf("Fail to remove lien [%s], error [%s]", name, err.Error())
		} else {
			logger.Printf("Removed lien [%s]", name)
		}
	}

	removeFirewallPolicies := func(folder string) {
		logger.Printf("Try to remove Firewall Policies from folder [%s]", folder)
		firewallPolicyList, err := firewallPoliciesService.List().ParentId(folder).Context(ctx).Do()
		if err != nil {
			logger.Printf("Fail to list Firewall Policies from folder [%s], error [%s]", folder, err.Error())
			return
		}
		for _, policy := range firewallPolicyList.Items {
			for _, association := range policy.Associations {
				_, err := firewallPoliciesService.RemoveAssociation(policy.Name).Name(association.Name).Context(ctx).Do()
				if err != nil {
					logger.Printf("Fail to Remove Association for Firewall Policies from folder [%s], error [%s]", folder, err.Error())
					return
				}
			}
			_, err := firewallPoliciesService.Delete(policy.Name).Context(ctx).Do()
			if err != nil {
				logger.Printf("Fail to delete Firewall Policy [%s] from folder [%s], error [%s]", policy.Name, folder, err.Error())
				return
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
			logger.Printf("Fail to list services for [%s], error [%s]", projectId, err.Error())
			return
		}

		if len(listResponse.Services) <= 1 {
			return
		}

		for _, service := range listResponse.Services {
			logger.Printf("Try to remove service: %s", service.ServiceName)
			_, err = endpointService.Services.Delete(service.ServiceName).Do()
			if err != nil {
				logger.Printf("Fail to delete service [%s] for [%s], error [%s]", service.ServiceName, projectId, err.Error())
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
			logger.Printf("Fail to remove project [%s], error [%s]", projectId, err.Error())
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
			logger.Printf("Fail to get all liens for the project [%s], error [%s]", projectId, err.Error())
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
			logger.Printf("Fail to get projects for the folder with id [%s], error [%s]", localFolderId, err.Error())
		} else {
			logger.Printf("Got and processed all projects for the folder with id [%s]", localFolderId)
		}
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
			if folder.Parent != fmt.Sprintf("folders/%s", rootFolderId) && folder.Name != fmt.Sprintf("folders/%s", rootFolderId) {
				removeFolder(folder)
			}
			return nil
		}); err != nil {
			logger.Fatalf("Fail to get subfolders for the folder with id [%s], error [%s]", folderId, err.Error())
		}
	}

	rootFolderId := fmt.Sprintf("folders/%s", rootFolderId)
	rootFolder, err := folderService.Get(rootFolderId).Do()
	if err != nil {
		logger.Printf("Fail to get parent folder [%s], error [%s]", rootFolderId, err.Error())
	} else {
		getSubFoldersAndRemoveProjectsFoldersRecursively(rootFolder, getSubFoldersAndRemoveProjectsFoldersRecursively)
	}
}

func CleanUpProjects(ctx context.Context, m PubSubMessage) error {
	invoke(ctx)
	return nil
}
