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

resource "random_pet" "main" {
  length    = 2
  separator = "-"
}

module "pubsub_scheduled_example" {
  source                    = "../../"
  project_id                = var.project_id
  job_name                  = "pubsub-example-${random_pet.main.id}"
  job_schedule              = "*/5 * * * *"
  function_entry_point      = "doSomething"
  function_source_directory = "${path.module}/function_source"
  function_name             = "testfunction-${random_pet.main.id}"
  region                    = var.region
  topic_name                = "pubsub_example_topic_${random_pet.main.id}"
}
