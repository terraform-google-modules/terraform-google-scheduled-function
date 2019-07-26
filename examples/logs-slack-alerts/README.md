# Logging Slack Alerts

This logging slack alerts example module schedules a job to run hourly queries of any errors which have occurred in logs which have been ingested into BigQuery. If any errors are found, the errors are sent as alerts to a slack webhook.

Running this module requires log exports into BigQuery in the specified project/region, which is not handled by this example. 
A good example of exported logging in BigQuery can be found in [Stackdriver Logging](https://cloud.google.com/logging/docs/export/).

[^]: (autogen_docs_start)

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| project\_id | The project ID to host the network in | string | n/a | yes |
| region | The region the project is in (App Engine specific) | string | `"us-central1"` | no |
| slack_webhook | The Slack webhook to send alerts | string | n/a | yes |
| dataset_name | The BigQuery Dataset where exported logging is sent | string | n/a | yes |
| audit_log_table | The BigQuery Table within the dataset where logging is sent | string | n/a | yes |
| time_column | The column within the BQ Table representing logging time | string | n/a | yes |
| error_message_column | The column within the BQ Table representing logging errors | string | n/a | yes |


[^]: (autogen_docs_end)