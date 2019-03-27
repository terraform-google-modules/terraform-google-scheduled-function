# Old Project Cleanup Utility Module

This example module schedules a job to clean up GCP projects older than a specified length of time, that match a particular key-value pair. This job runs every 5 minutes via Google Cloud Scheduled Functions. Please see the [utility's readme](./function_source/README.md) for more information as to its operation and configuration.

Running this module requires an App Engine app in the specified project/region, which is not handled by this example. More information is in the [root readme](../../README.md#app-engine).

[^]: (autogen_docs_start)


## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| organization_id | The organization ID whose projects to clean up | string | - | yes |
| project_id | The project ID to host the scheduled function in | string | - | yes |
| region | The region the project is in (App Engine specific) | string | `us-central1` | no |

## Outputs

| Name | Description |
|------|-------------|
| name | The name of the job created |
| project_id | The project ID |

[^]: (autogen_docs_end)
