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
  version             = "~> 8.0"
  topic               = var.topic_name
  project_id          = var.project_id
  create_topic        = var.scheduler_job == null ? true : false
  grant_token_creator = var.grant_token_creator
  topic_labels        = var.topic_labels
  topic_kms_key_name  = var.topic_kms_key_name
}

/******************************************
	Cloud Function Resource Definitions
 *****************************************/

resource "random_id" "suffix" {
  byte_length = 4
}

data "archive_file" "main" {
  type        = "zip"
  source_dir  = var.function_source_directory
  output_path = "${path.module}/.temp/${var.function_name}.zip"
  excludes    = concat([".gcloudignore", "README.md", ".git", ".gitignore"], var.files_to_exclude_in_source_dir)
  depends_on  = [var.function_source_dependent_files]
}

resource "google_storage_bucket" "main" {
  count                       = var.create_bucket ? 1 : 0
  name                        = var.bucket_name == "" ? "${var.project_id}-scheduled-function-${random_id.suffix.hex}" : var.bucket_name
  project                     = var.project_id
  location                    = var.region
  force_destroy               = var.bucket_force_destroy
  uniform_bucket_level_access = true
  labels                      = var.function_source_archive_bucket_labels
}

resource "google_storage_bucket_object" "main" {
  name   = "${var.function_name}-${data.archive_file.main.output_md5}.zip"
  bucket = var.create_bucket ? google_storage_bucket.main[0].name : var.bucket_name
  source = data.archive_file.main.output_path
}

resource "google_cloudfunctions2_function" "main" {
  name         = var.function_name
  location     = var.region
  project      = var.project_id
  description  = var.function_description
  labels       = var.function_labels
  kms_key_name = var.function_kms_key_name

  build_config {
    runtime           = var.function_runtime
    entry_point       = var.function_entry_point
    docker_repository = var.function_docker_repository != null ? var.function_docker_repository : (var.function_docker_registry != null ? null : null)
    source {
      storage_source {
        bucket = var.create_bucket ? google_storage_bucket.main[0].name : var.bucket_name
        object = google_storage_bucket_object.main.name
      }
    }
  }

  service_config {
    max_instance_count            = var.function_max_instances
    available_memory              = "${var.function_available_memory_mb}Mi"
    timeout_seconds               = var.function_timeout_s
    service_account_email         = var.function_service_account_email
    environment_variables         = var.function_environment_variables
    ingress_settings              = var.ingress_settings
    vpc_connector                 = var.vpc_connector
    vpc_connector_egress_settings = var.vpc_connector_egress_settings

    dynamic "secret_environment_variables" {
      for_each = var.function_secret_environment_variables
      content {
        key        = secret_environment_variables.value["key"]
        project_id = secret_environment_variables.value["project_id"]
        secret     = secret_environment_variables.value["secret_name"]
        version    = secret_environment_variables.value["version"]
      }
    }
  }

  event_trigger {
    trigger_region = var.region
    event_type     = "google.cloud.pubsub.topic.v1.messagePublished"
    pubsub_topic   = var.scheduler_job == null ? module.pubsub_topic.id : "projects/${var.project_id}/topics/${var.topic_name}"
    retry_policy   = var.function_event_trigger_failure_policy_retry ? "RETRY_POLICY_RETRY" : "RETRY_POLICY_DO_NOT_RETRY"
  }
}

