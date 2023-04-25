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
  type        = string
  description = "The ID of the project where the resources will be created"
}

variable "job_name" {
  type        = string
  description = "The name of the scheduled job to run"
  default     = null
}

variable "job_description" {
  type        = string
  description = "Addition text to describe the job"
  default     = ""
}

variable "job_schedule" {
  type        = string
  description = "The job frequency, in cron syntax"
  default     = "*/2 * * * *"
}

variable "function_available_memory_mb" {
  type        = number
  default     = 256
  description = "The amount of memory in megabytes allotted for the function to use."
}

variable "function_description" {
  type        = string
  default     = "Processes log export events provided through a Pub/Sub topic subscription."
  description = "The description of the function."
}

variable "function_entry_point" {
  type        = string
  description = "The name of a method in the function source which will be invoked when the function is executed."
}

variable "function_environment_variables" {
  type        = map(string)
  default     = {}
  description = "A set of key/value environment variable pairs to assign to the function."
}

variable "function_secret_environment_variables" {
  type        = list(map(string))
  default     = []
  description = "A list of maps which contains key, project_id, secret_name (not the full secret id) and version to assign to the function as a set of secret environment variables."
}

variable "function_event_trigger_failure_policy_retry" {
  type        = bool
  default     = false
  description = "A toggle to determine if the function should be retried on failure."
}

variable "function_labels" {
  type        = map(string)
  default     = {}
  description = "A set of key/value label pairs to assign to the function."
}

variable "function_runtime" {
  type        = string
  default     = "nodejs10"
  description = "The runtime in which the function will be executed."
}

variable "function_source_archive_bucket_labels" {
  type        = map(string)
  default     = {}
  description = "A set of key/value label pairs to assign to the function source archive bucket."
}

variable "function_source_directory" {
  type        = string
  description = "The contents of this directory will be archived and used as the function source."
}

variable "function_source_dependent_files" {
  type = list(object({
    filename = string
    id       = string
  }))
  description = "A list of any terraform created `local_file`s that the module will wait for before creating the archive."
  default     = []
}

variable "function_timeout_s" {
  type        = number
  default     = 60
  description = "The amount of time in seconds allotted for the execution of the function."
}

variable "function_service_account_email" {
  type        = string
  default     = ""
  description = "The service account to run the function as."
}

variable "function_max_instances" {
  type        = number
  default     = null
  description = "The maximum number of parallel executions of the function."
}

variable "ingress_settings" {
  type        = string
  default     = null
  description = "The ingress settings for the function. Allowed values are ALLOW_ALL, ALLOW_INTERNAL_AND_GCLB and ALLOW_INTERNAL_ONLY. Changes to this field will recreate the cloud function."
}

variable "function_docker_registry" {
  type        = string
  default     = "CONTAINER_REGISTRY"
  description = "Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY."
}

variable "function_docker_repository" {
  type        = string
  default     = ""
  description = "User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry."
}

variable "function_kms_key_name" {
  type        = string
  default     = ""
  description = "Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources."
}

variable "vpc_connector" {
  type        = string
  default     = null
  description = "The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is projects//locations//connectors/*."
}

variable "vpc_connector_egress_settings" {
  type        = string
  default     = null
  description = "The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are ALL_TRAFFIC and PRIVATE_RANGES_ONLY. If unset, this field preserves the previously set value."
}

variable "bucket_name" {
  type        = string
  default     = ""
  description = "The name to apply to the bucket. Will default to a string of <project-id>-scheduled-function-XXXX> with XXXX being random characters."
}

variable "bucket_force_destroy" {
  type        = bool
  default     = true
  description = "When deleting the GCS bucket containing the cloud function, delete all objects in the bucket first."
}

variable "function_name" {
  type        = string
  description = "The name to apply to the function"
}

variable "region" {
  type        = string
  description = "The region in which resources will be applied."
}

variable "topic_name" {
  type        = string
  description = "Name of pubsub topic connecting the scheduled job and the function"
  default     = "test-topic"
}

variable "topic_labels" {
  type        = map(string)
  description = "A set of key/value label pairs to assign to the pubsub topic."
  default     = {}
}

variable "message_data" {
  type        = string
  description = "The data to send in the topic message."
  default     = "dGVzdA=="
}

variable "time_zone" {
  type        = string
  description = "The timezone to use in scheduler"
  default     = "Etc/UTC"
}

variable "scheduler_job" {
  type        = object({ name = string })
  description = "An existing Cloud Scheduler job instance"
  default     = null
}

variable "grant_token_creator" {
  type        = bool
  description = "Specify true if you want to add token creator role to the default Pub/Sub SA"
  default     = false
}
