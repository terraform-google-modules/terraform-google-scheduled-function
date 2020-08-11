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

provider "google-beta" {
  version = "~> 3.33"
  project = var.project_id
  region  = var.region
}

provider "google" {
  version = "~> 3.33"
  project = var.project_id
  region  = var.region
}

module "log_slack_alerts_example" {
  providers = {
    google = google-beta
  }

  source                    = "../../"
  project_id                = var.project_id
  job_name                  = "logs_query"
  job_description           = "Scheduled time to run audit query to check for errors"
  job_schedule              = "55 * * * *"
  function_entry_point      = "query_for_errors"
  function_source_directory = "${path.module}/function_source"
  function_name             = "logs_query_alerting"
  function_description      = "Cloud Function to query audit logs for errors"
  region                    = var.region
  topic_name                = "logs_query_topic"
  function_runtime          = "python37"

  function_environment_variables = {
    SLACK_WEBHOOK        = var.slack_webhook
    DATASET_NAME         = var.dataset_name
    AUDIT_LOG_TABLE      = var.audit_log_table
    TIME_COLUMN          = var.time_column
    ERROR_MESSAGE_COLUMN = var.error_message_column
  }
}

