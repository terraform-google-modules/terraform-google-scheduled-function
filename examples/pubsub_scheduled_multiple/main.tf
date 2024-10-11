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

module "pubsub_scheduled_1" {
  source  = "terraform-google-modules/scheduled-function/google"
  version = "~> 6.0"

  project_id                = var.project_id
  job_name                  = "pubsub-example"
  job_schedule              = "*/5 * * * *"
  function_entry_point      = "doSomething"
  function_source_directory = "${path.module}/function_source"
  function_name             = "testfunction-1"
  region                    = var.region
  topic_name                = "pubsub-1"
}

module "pubsub_scheduled_2" {
  source  = "terraform-google-modules/scheduled-function/google"
  version = "~> 6.0"

  project_id                = var.project_id
  function_entry_point      = "doSomething2"
  function_source_directory = "${path.module}/function_source_2"
  function_name             = "testfunction-2"
  region                    = var.region
  topic_name                = module.pubsub_scheduled_1.pubsub_topic_name
  scheduler_job             = module.pubsub_scheduled_1.scheduler_job
}
