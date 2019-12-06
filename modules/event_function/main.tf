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


/******************************************
	Cloud Function Resource Definitions
 *****************************************/

module "main" {
  source  = "terraform-google-modules/event-function/google"
  version = "~> 1.1"

  entry_point = var.function_entry_point
  event_trigger = {
    event_type = "google.pubsub.topic.publish"
    resource   = var.topic_name
  }
  name             = var.function_name
  project_id       = var.project_id
  region           = var.region
  runtime          = var.function_runtime
  source_directory = var.function_source_directory

  available_memory_mb                = var.function_available_memory_mb
  bucket_force_destroy               = var.bucket_force_destroy
  bucket_labels                      = var.function_source_archive_bucket_labels
  bucket_name                        = var.bucket_name
  description                        = var.function_description
  environment_variables              = var.function_environment_variables
  event_trigger_failure_policy_retry = var.function_event_trigger_failure_policy_retry
  labels                             = var.function_labels
  service_account_email              = var.function_service_account_email
  timeout_s                          = var.function_timeout_s
}
