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

variable "function_timeout_s" {
  type        = number
  default     = 500
  description = "The amount of time in seconds allotted for the execution of the function."
}

variable "organization_id" {
  type        = string
  description = "The organization ID whose projects to clean up"
}

variable "clean_up_org_level_cai_feeds" {
  type        = bool
  description = "Clean up organization level Cloud Asset Inventory Feeds."
  default     = false
}

variable "target_included_feeds" {
  type        = list(string)
  description = "List of organization level Cloud Asset Inventory feeds that should be deleted. Regex example: `.*/feeds/fd-cai-monitoring-.*` "
  default     = []
}

variable "project_id" {
  type        = string
  description = "The project ID to host the scheduled function in"
}

variable "region" {
  type        = string
  description = "The region the project is in (App Engine specific)"
}

variable "job_schedule" {
  type        = string
  description = "Cleaner function run frequency, in cron syntax"
  default     = "*/5 * * * *"
}

variable "topic_name" {
  type        = string
  description = "Name of pubsub topic connecting the scheduled projects cleanup function"
  default     = "pubsub_scheduled_project_cleaner"
}

variable "target_tag_name" {
  type        = string
  description = "The name of a tag to filter GCP projects on for consideration by the cleanup utility (legacy, use `target_included_labels` map instead)."
  default     = ""
}

variable "target_tag_value" {
  type        = string
  description = "The value of a tag to filter GCP projects on for consideration by the cleanup utility (legacy, use `target_included_labels` map instead)."
  default     = ""
}

variable "max_project_age_in_hours" {
  type        = number
  description = "The maximum number of hours that a GCP project, selected by `target_tag_name` and `target_tag_value`, can exist"
  default     = 6
}

variable "target_excluded_labels" {
  type        = map(string)
  description = "Map of project lablels that won't be deleted."
  default     = {}
}

variable "target_included_labels" {
  type        = map(string)
  description = "Map of project lablels that will be deleted."
  default     = {}
}

variable "clean_up_org_level_scc_notifications" {
  type        = bool
  description = "Clean up organization level Security Command Center notifications."
  default     = false
}

variable "target_included_scc_notifications" {
  type        = list(string)
  description = "List of organization Security Command Center notifications names regex that will be deleted. Regex example: `.*/notificationConfigs/scc-notify-.*` "
  default     = []
}

variable "list_scc_notifications_page_size" {
  type        = number
  description = "The maximum number of notification configs to return in the call to `ListNotificationConfigs` service. The minimun value is 1 and the maximum value is 1000."
  default     = 500
}

variable "clean_up_org_level_tag_keys" {
  type        = bool
  description = "Clean up organization level Tag Keys."
  default     = false
}

variable "target_excluded_tagkeys" {
  type        = list(string)
  description = "List of organization Tag Key short names that won't be deleted."
  default     = []
}

variable "clean_up_billing_sinks" {
  type        = bool
  description = "Clean up Billing Account Sinks."
  default     = false
}

variable "billing_account" {
  type        = string
  description = "Billing Account used to provision resources."
  default     = ""
}

variable "target_billing_sinks" {
  type        = list(string)
  description = "List of Billing Account Log Sinks names regex that will be deleted. Regex example: `.*/sinks/sk-c-logging-.*-billing-.*` "
  default     = []
}

variable "list_billing_sinks_page_size" {
  type        = number
  description = "The maximum number of Billing Account Log Sinks to return in the call to `BillingAccountsSinksService.List` service."
  default     = 200
}

variable "target_folder_id" {
  type        = string
  description = "Folder ID to delete all projects under."
  default     = ""
}

variable "function_docker_registry" {
  type        = string
  default     = null
  description = "Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY."
}
