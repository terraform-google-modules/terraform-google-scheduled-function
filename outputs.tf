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

output "name" {
  value       = length(google_cloud_scheduler_job.job) > 0 ? google_cloud_scheduler_job.job.0.name : null
  description = "The name of the job created"
}

output "scheduler_job" {
  value       = length(google_cloud_scheduler_job.job) > 0 ? google_cloud_scheduler_job.job.0 : null
  description = "The Cloud Scheduler job instance"
}

output "pubsub_topic_name" {
  value       = var.scheduler_job == null ? module.pubsub_topic.topic : var.topic_name
  description = "PubSub topic name"
}
