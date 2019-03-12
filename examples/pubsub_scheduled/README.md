# Simple Project

This example module schedules a job to publish a message to a Pub/Sub topic every 5 minutes, which will trigger a CloudFunctions function.

Note that this example requires an app_engine app in the specified project/region. This is because scheduled functions are dependent on the project being configured with app engine.

Documentation on the app engine dependency is here: https://cloud.google.com/scheduler/docs/ 

[^]: (autogen_docs_start)

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

[^]: (autogen_docs_end)
