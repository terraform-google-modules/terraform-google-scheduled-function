# Simple Project

This example module schedules a job to publish a message to a Pub/Sub topic every 5 minutes, which will trigger a CloudFunctions function.

Running this module requires an App Engine app in the specified project/region, which is not handled by this example.
More information is in the [root readme](../../README.md#app-engine).

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| project\_id | The project ID to host the network in | string | n/a | yes |
| region | The region the project is in (App Engine specific) | string | `"us-central1"` | no |

## Outputs

| Name | Description |
|------|-------------|
| name | The name of the job created |
| project\_id | The project ID |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
