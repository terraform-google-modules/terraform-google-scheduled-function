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
	Scheduled Function Definition
 *****************************************/

resource "google_cloud_scheduler_job" "job" {
  count = var.scheduler_job == null ? 1 : 0

  name        = var.job_name
  project     = var.project_id
  region      = var.region
  description = var.job_description
  schedule    = var.job_schedule
  time_zone   = var.time_zone

  pubsub_target {
    topic_name = "projects/${var.project_id}/topics/${module.pubsub_topic.topic}"
    data       = var.message_data
  }
}

/******************************************
	PubSub Topic Definition
 *****************************************/

module "pubsub_topic" {
  source              = "terraform-google-modules/pubsub/google"
  version             = "~> 1.0"
  topic               = var.topic_name
  project_id          = var.project_id
  create_topic        = var.scheduler_job == null ? true : false
  grant_token_creator = var.grant_token_creator
}

/******************************************
	Cloud Function Resource Definitions
 *****************************************/

resource "random_id" "suffix" {

  byte_length = 4
}

module "main" {
  source  = "terraform-google-modules/event-function/google"
  version = "~> 1.6"

  entry_point = var.function_entry_point
  event_trigger = {
    event_type = "google.pubsub.topic.publish"
    resource   = var.scheduler_job == null ? module.pubsub_topic.topic : var.topic_name
  }
  name             = var.function_name
  project_id       = var.project_id
  region           = var.region
  runtime          = var.function_runtime
  source_directory = var.function_source_directory

  source_dependent_files = var.function_source_dependent_files

  available_memory_mb                = var.function_available_memory_mb
  bucket_force_destroy               = var.bucket_force_destroy
  bucket_labels                      = var.function_source_archive_bucket_labels
  bucket_name                        = var.bucket_name == "" ? "${var.project_id}-scheduled-function-${random_id.suffix.hex}" : var.bucket_name
  description                        = var.function_description
  environment_variables              = var.function_environment_variables
  event_trigger_failure_policy_retry = var.function_event_trigger_failure_policy_retry
  labels                             = var.function_labels
  service_account_email              = var.function_service_account_email
  timeout_s                          = var.function_timeout_s
  vpc_connector                      = var.function_vpc_connector
  vpc_connector_egress_settings      = var.function_vpc_connector_egress_settings
}
