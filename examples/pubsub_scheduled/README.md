# Simple Project

This example module schedules a job to publish a message to a Pub/Sub topic every 5 minutes, which will trigger a CloudFunctions function.

[^]: (autogen_docs_start)

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| function\_source | The directory of the source code for the cloud functions function | string | `"./function_source"` | no |
| project\_id | The project ID to host the network in | string | n/a | yes |
| region | The region the project is in (app engine specific) | string | `"us-central1-b"` | no |

## Outputs

| Name | Description |
|------|-------------|
| name | The name of the job created |
| project\_id | The prject id |

[^]: (autogen_docs_end)
