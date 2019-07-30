/**
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

resource "google_service_account" "project_cleaner_function" {
  project      = var.project_id
  account_id   = "project-cleaner-function"
  display_name = "Project Cleaner Function"
}

resource "google_organization_iam_member" "project_owner" {
  org_id = var.organization_id
  role   = "roles/owner"
  member = "serviceAccount:${google_service_account.project_cleaner_function.email}"
}

module "scheduled_project_cleaner" {
  source                         = "../../"
  project_id                     = var.project_id
  job_name                       = "project-cleaner"
  job_schedule                   = "*/5 * * * *"
  function_entry_point           = "CleanUpProjects"
  function_source_directory      = "${path.module}/function_source"
  function_name                  = "old-project-cleaner"
  region                         = var.region
  topic_name                     = "pubsub_scheduled_project_cleaner"
  function_available_memory_mb   = 128
  function_description           = "Clean up GCP projects older than ${var.max_project_age_in_hours} hours matching particular tags"
  function_runtime               = "go111"
  function_service_account_email = "${google_service_account.project_cleaner_function.email}"

  function_environment_variables = {
    TARGET_TAG_NAME       = var.target_tag_name
    TARGET_TAG_VALUE      = var.target_tag_value
    MAX_PROJECT_AGE_HOURS = var.max_project_age_in_hours
  }
}
