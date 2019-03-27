// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package project_cleaner

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

// PubSubMessage wraps the message sent to the background Cloud Function by GCP PubSub.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// CleanUpProjects is the entry point of the scheduled function.
func CleanUpProjects(ctx context.Context, m PubSubMessage) error {
	targetTag := os.Getenv("TARGET_TAG_NAME")
	targetValue := os.Getenv("TARGET_TAG_VALUE")
	maxAgeInHoursStr := os.Getenv("MAX_PROJECT_AGE_HOURS")
	maxAgeInHours, err := strconv.ParseInt(maxAgeInHoursStr, 10, 0)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Could not convert %s to integer", maxAgeInHoursStr))
		return err
	}
	logger.Println(fmt.Sprintf("Launching project cleanup for projects older than %d hours, with tag %s=%s", maxAgeInHours, targetTag, targetValue))
	err = deleteProjectsMatchingTag(ctx, targetTag, targetValue, maxAgeInHours)
	if err != nil {
		logger.Fatal(err)
	}
	return err
}

func deleteProjectsMatchingTag(ctx context.Context, key string, value string, acceptableAgeInHours int64) error {
	logger.Println("Initializing Google client")
	c, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		return err
	}

	cloudResourceManagerService, err := cloudresourcemanager.New(c)
	if err != nil {
		return err
	}

	resourceCreationCutoff := tooOldTime(int64(acceptableAgeInHours * 60 * 60))

	logger.Println("Looking through projects")
	req := cloudResourceManagerService.Projects.List()
	if err := req.Pages(ctx, func(page *cloudresourcemanager.ListProjectsResponse) error {
		for _, project := range page.Projects {
			if val, ok := project.Labels[key]; ok && val == value {
				logger.Println(fmt.Sprintf("Considering project %s...", project.ProjectId))
				if project.LifecycleState == "ACTIVE" {
					projectCreatedAt, err := time.Parse(time.RFC3339, project.CreateTime)
					if err != nil {
						return err
					}
					if projectCreatedAt.Before(resourceCreationCutoff) {
						logger.Println(fmt.Sprintf("Project %s was created at %s, is active, and is older than %d hours. Deleting.", project.ProjectId, project.CreateTime, acceptableAgeInHours))
						_, err = cloudResourceManagerService.Projects.Delete(project.ProjectId).Do()
						if err != nil {
							logger.Fatal(err)
							return err
						}
						logger.Println(fmt.Sprintf("Requested deletion of project %s.", project.ProjectId))
					}
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	logger.Println("Considered all projects")
	return nil
}

func tooOldTime(i int64) time.Time {
	return time.Unix(time.Now().Unix()-i, 0)
}
