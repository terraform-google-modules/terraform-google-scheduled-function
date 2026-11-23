# Multiple Cloud Functions sharing the same Cloud Scheduler

This example module schedules a job to publish a message to a Pub/Sub topic every 5 minutes, which will trigger a CloudFunctions function.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| project\_id | The project ID to host the network in | `string` | `"flask-app-254610"` | no |
| region | The region the project is in | `string` | `"us-central1"` | no |

## Outputs

| Name | Description |
|------|-------------|
| name | The name of the job created |
| project\_id | The project ID |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
