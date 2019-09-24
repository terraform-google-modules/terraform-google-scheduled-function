# Logging Slack Alerts

Disclaimer (8/1/2019): Test Coverage has currently not been added to this example and may be unintentionally broken by future releases to the module.

This logging slack alerts example module schedules a job to run hourly queries of any errors which have occurred in logs which have been ingested into BigQuery. If any errors are found, the errors are sent as alerts to a slack webhook.

Running this module requires log exports into BigQuery in the specified project/region, which is not handled by this example. 
A good example of exported logging in BigQuery can be found in [Stackdriver Logging](https://cloud.google.com/logging/docs/export/).

## Configure a Service Account

If not using the default App Engine default service account (PROJECT_ID@appspot.gserviceaccount.com), which has the Editor role on the project, one can configure a service account for the cloud function which has the following IAM role - (roles/bigquery.dataViewer, roles/bigquery.jobUser). Additionally, the BigQuery API (https://bigquery.googleapis.com) needs to be enabled as well.


<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| audit\_log\_table | BigQuery Table where logs are sent | string | n/a | yes |
| dataset\_name | BigQuery Dataset where logs are sent | string | n/a | yes |
| error\_message\_column | BigQuery Column in audit log table representing logging error | string | n/a | yes |
| job\_schedule | The cron schedule for triggering the cloud function | string | `"55 * * * *"` | no |
| project\_id | The project ID to host the network in | string | n/a | yes |
| region | The region the project is in (App Engine specific) | string | `"us-central1"` | no |
| slack\_webhook | Slack webhook to send alerts | string | n/a | yes |
| time\_column | BigQuery Column in audit log table representing logging time | string | n/a | yes |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
