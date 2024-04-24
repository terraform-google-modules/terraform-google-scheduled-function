# Old Project Cleanup Utility Module

This module schedules a job to clean up GCP projects older than a specified length of time, that match a particular labels. This job runs every 5 minutes via Google Cloud Scheduled Functions. Please see the [utility's readme](./function_source/README.md) for more information as to its operation and configuration.

## Requirements

### App Engine

Running this module requires an App Engine app in the specified project/region. More information is in the [root readme](../../README.md#app-engine).

### Enabled Services

The following services must be enabled on the project housing the cleanup function prior to invoking this module:

- Cloud Functions (`cloudfunctions.googleapis.com`)
- Cloud Scheduler (`cloudscheduler.googleapis.com`)
- Cloud Resource Manager (`cloudresourcemanager.googleapis.com`)
- Compute Engine API (`compute.googleapis.com`)
- Security Command Center API (`securitycenter.googleapis.com`)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| clean\_up\_org\_level\_scc\_notifications | Clean up organization level Security Command Center notifications. | `bool` | `false` | no |
| clean\_up\_org\_level\_tag\_keys | Clean up organization level Tag Keys. | `bool` | `false` | no |
| function\_timeout\_s | The amount of time in seconds allotted for the execution of the function. | `number` | `500` | no |
| job\_schedule | Cleaner function run frequency, in cron syntax | `string` | `"*/5 * * * *"` | no |
| list\_scc\_notifications\_page\_size | The maximum number of notification configs to return in the call to `ListNotificationConfigs` service. The minimun value is 1 and the maximum value is 1000. | `number` | `500` | no |
| max\_project\_age\_in\_hours | The maximum number of hours that a GCP project, selected by `target_tag_name` and `target_tag_value`, can exist | `number` | `6` | no |
| organization\_id | The organization ID whose projects to clean up | `string` | n/a | yes |
| project\_id | The project ID to host the scheduled function in | `string` | n/a | yes |
| region | The region the project is in (App Engine specific) | `string` | n/a | yes |
| target\_excluded\_labels | Map of project lablels that won't be deleted. | `map(string)` | `{}` | no |
| target\_excluded\_tagkeys | List of organization Tag Key short names that won't be deleted. | `list(string)` | `[]` | no |
| target\_folder\_id | Folder ID to delete all projects under. | `string` | `""` | no |
| target\_included\_labels | Map of project lablels that will be deleted. | `map(string)` | `{}` | no |
| target\_included\_scc\_notifications | List of organization  Security Command Center notifications names regex that will be deleted. Regex example: `.*/notificationConfigs/scc-notify-.*` | `list(string)` | `[]` | no |
| target\_tag\_name | The name of a tag to filter GCP projects on for consideration by the cleanup utility (legacy, use `target_included_labels` map instead). | `string` | `""` | no |
| target\_tag\_value | The value of a tag to filter GCP projects on for consideration by the cleanup utility (legacy, use `target_included_labels` map instead). | `string` | `""` | no |
| topic\_name | Name of pubsub topic connecting the scheduled projects cleanup function | `string` | `"pubsub_scheduled_project_cleaner"` | no |

## Outputs

| Name | Description |
|------|-------------|
| name | The name of the job created |
| project\_id | The project ID |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
