# Simple Project

This example module schedules a job to publish a message to a Pub/Sub topic every 5 minutes, which will trigger a CloudFunctions function.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| project\_id | The project ID to host the network in | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| function\_name | The name of the Cloud Function |
| name | The name of the job created |
| project\_id | The project ID |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
