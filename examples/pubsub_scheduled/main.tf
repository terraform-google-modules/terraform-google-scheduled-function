/**
 * Copyright 2018 Google LLC
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

module "pubsub_scheduled_example" {
  source                    = "../../"
  project_id                = "${var.project_id}"
  job_name                  = "pubsub-example"
  job_schedule              = "*/5 * * * *"
  function_entry_point      = "doSomething"
  function_source_directory = "${var.function_source}"
  name                      = "testfunction-foo"
  region                    = "us-central1"
  topic_name                = "pubsub_example_topic"
}
