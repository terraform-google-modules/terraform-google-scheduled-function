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

variable "organization_id" {
  type        = string
  description = "The organization ID whose projects to clean up"
}

variable "project_id" {
  type        = string
  description = "The project ID to host the scheduled function in"
}

variable "region" {
  type        = string
  description = "The region the project is in (App Engine specific)"
}

variable "target_tag_name" {
  type        = string
  description = "The name of a tag to filter GCP projects on for consideration by the cleanup utility"
  default     = "cft-ephemeral"
}

variable "target_tag_value" {
  type        = string
  description = "The value of a tag to filter GCP projects on for consideration by the cleanup utility"
  default     = "true"
}

variable "max_project_age_in_hours" {
  type        = number
  description = "The maximum number of hours that a GCP project, selected by `target_tag_name` and `target_tag_value`, can exist"
  default     = 6
}
