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

variable "project_id" {
  description = "The project ID to host the network in"
  type        = string
}

variable "slack_webhook" {
  description = "Slack webhook to send alerts"
  type        = string
}

variable "dataset_name" {
  description = "BigQuery Dataset where logs are sent"
  type        = string
}

variable "audit_log_table" {
  description = "BigQuery Table where logs are sent"
  type        = string
}

variable "time_column" {
  description = "BigQuery Column in audit log table representing logging time"
  type        = string
}

variable "error_message_column" {
  description = "BigQuery Column in audit log table representing logging error"
  type        = string
}

variable "region" {
  description = "The region the project is in (App Engine specific)"
  type        = string
  default     = "us-central1"
}

