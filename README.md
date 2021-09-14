# Terraform Google Scheduled Functions Module

This modules makes it easy to set up a scheduled job to trigger events/run functions.

## Compatibility
This module is meant for use with Terraform 0.13. If you haven't
[upgraded](https://www.terraform.io/upgrade-guides/0-13.html) and need a Terraform
0.12.x-compatible version of this module, the last released version
intended for Terraform 0.12.x is [v1.6.0](https://registry.terraform.io/modules/terraform-google-modules/scheduled-function/google/1.6.0).

## Usage
You can go to the examples folder, however the usage of the module could be like this in your own main.tf file:

```hcl
module "scheduled-function" {
  source  = "terraform-google-modules/scheduled-function/google"
  version = "0.1.0"
  project_id   = "<PROJECT ID>"
  job_name="<NAME_OF_JOB>"
  job_schedule="<CRON_SYNTAX_SCHEDULE>"
  function_entry_point="<NAME_OF_FUNCTION>"
  function_source_directory="<DIRECTORY_OF_FUNCTION_SOURCE>"
  function_name="<RESOURCE_NAMES>"
  region="<REGION>"
}
```

Then perform the following commands on the root folder:

- `terraform init` to get the plugins
- `terraform plan` to see the infrastructure plan
- `terraform apply` to apply the infrastructure build
- `terraform destroy` to destroy the built infrastructure

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_bucket_force_destroy"></a> [bucket\_force\_destroy](#input\_bucket\_force\_destroy) | When deleting the GCS bucket containing the cloud function, delete all objects in the bucket first. | `bool` | `true` | no |
| <a name="input_bucket_name"></a> [bucket\_name](#input\_bucket\_name) | The name to apply to the bucket. Will default to a string of `<project-id>-scheduled-function-<XXXX>` with `XXXX` being random characters. | `string` | `""` | no |
| <a name="input_function_available_memory_mb"></a> [function\_available\_memory\_mb](#input\_function\_available\_memory\_mb) | The amount of memory in megabytes allotted for the function to use. | `number` | `256` | no |
| <a name="input_function_description"></a> [function\_description](#input\_function\_description) | The description of the function. | `string` | `"Processes log export events provided through a Pub/Sub topic subscription."` | no |
| <a name="input_function_entry_point"></a> [function\_entry\_point](#input\_function\_entry\_point) | The name of a method in the function source which will be invoked when the function is executed. | `string` | n/a | yes |
| <a name="input_function_environment_variables"></a> [function\_environment\_variables](#input\_function\_environment\_variables) | A set of key/value environment variable pairs to assign to the function. | `map(string)` | `{}` | no |
| <a name="input_function_event_trigger_failure_policy_retry"></a> [function\_event\_trigger\_failure\_policy\_retry](#input\_function\_event\_trigger\_failure\_policy\_retry) | A toggle to determine if the function should be retried on failure. | `bool` | `false` | no |
| <a name="input_function_labels"></a> [function\_labels](#input\_function\_labels) | A set of key/value label pairs to assign to the function. | `map(string)` | `{}` | no |
| <a name="input_function_name"></a> [function\_name](#input\_function\_name) | The name to apply to the function | `string` | n/a | yes |
| <a name="input_function_runtime"></a> [function\_runtime](#input\_function\_runtime) | The runtime in which the function will be executed. | `string` | `"nodejs10"` | no |
| <a name="input_function_service_account_email"></a> [function\_service\_account\_email](#input\_function\_service\_account\_email) | The service account to run the function as. | `string` | `""` | no |
| <a name="input_function_source_archive_bucket_labels"></a> [function\_source\_archive\_bucket\_labels](#input\_function\_source\_archive\_bucket\_labels) | A set of key/value label pairs to assign to the function source archive bucket. | `map(string)` | `{}` | no |
| <a name="input_function_source_dependent_files"></a> [function\_source\_dependent\_files](#input\_function\_source\_dependent\_files) | A list of any terraform created `local_file`s that the module will wait for before creating the archive. | <pre>list(object({<br>    filename = string<br>    id       = string<br>  }))</pre> | `[]` | no |
| <a name="input_function_source_directory"></a> [function\_source\_directory](#input\_function\_source\_directory) | The contents of this directory will be archived and used as the function source. | `string` | n/a | yes |
| <a name="input_function_timeout_s"></a> [function\_timeout\_s](#input\_function\_timeout\_s) | The amount of time in seconds allotted for the execution of the function. | `number` | `60` | no |
| <a name="input_grant_token_creator"></a> [grant\_token\_creator](#input\_grant\_token\_creator) | Specify true if you want to add token creator role to the default Pub/Sub SA | `bool` | `false` | no |
| <a name="input_job_description"></a> [job\_description](#input\_job\_description) | Addition text to describe the job | `string` | `""` | no |
| <a name="input_job_name"></a> [job\_name](#input\_job\_name) | The name of the scheduled job to run | `string` | `null` | no |
| <a name="input_job_schedule"></a> [job\_schedule](#input\_job\_schedule) | The job frequency, in 5-field, unix-cron syntax (excluding seconds) | `string` | `"*/2 * * * *"` | no |
| <a name="input_message_data"></a> [message\_data](#input\_message\_data) | The data to send in the topic message. | `string` | `"dGVzdA=="` | no |
| <a name="input_project_id"></a> [project\_id](#input\_project\_id) | The ID of the project where the resources will be created | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | The region in which resources will be applied. | `string` | n/a | yes |
| <a name="input_scheduler_job"></a> [scheduler\_job](#input\_scheduler\_job) | An existing Cloud Scheduler job instance | `object({ name = string })` | `null` | no |
| <a name="input_time_zone"></a> [time\_zone](#input\_time\_zone) | The timezone to use in scheduler | `string` | `"Etc/UTC"` | no |
| <a name="input_topic_name"></a> [topic\_name](#input\_topic\_name) | Name of pubsub topic connecting the scheduled job and the function | `string` | `"test-topic"` | no |
| <a name="input_vpc_connector"></a> [vpc\_connector](#input\_vpc\_connector) | The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is projects//locations//connectors/*. | `string` | `null` | no |
| <a name="input_vpc_connector_egress_settings"></a> [vpc\_connector\_egress\_settings](#input\_vpc\_connector\_egress\_settings) | The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are ALL\_TRAFFIC and PRIVATE\_RANGES\_ONLY. If unset, this field preserves the previously set value. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_name"></a> [name](#output\_name) | The name of the job created |
| <a name="output_pubsub_topic_name"></a> [pubsub\_topic\_name](#output\_pubsub\_topic\_name) | PubSub topic name |
| <a name="output_scheduler_job"></a> [scheduler\_job](#output\_scheduler\_job) | The Cloud Scheduler job instance |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

These sections describe requirements for using this module.

### Software

The following dependencies must be available:

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.0
- [Terraform Provider for GCP][terraform-provider-gcp] plugin v2.14

### App Engine
Note that this module requires App Engine being configured in the specified project/region.
This is because Google Cloud Scheduler is dependent on the project being configured with App Engine.
Refer to the [Google Cloud Scheduler documentation][cloud-scheduler-documentation]
information on the App Engine dependency.

The recommended way to create projects with App Engine enabled is via the [Project Factory module](https://github.com/terraform-google-modules/terraform-google-project-factory).
There is an example of how to create the project [within that module](https://github.com/terraform-google-modules/terraform-google-project-factory/tree/master/examples/app_engine)

### Service Account

A service account with the following roles must be used to provision
the resources of this module:

- Storage Admin: `roles/storage.admin`
- PubSub Editor: `roles/pubsub.editor`
- Cloudscheduler Admin: `roles/cloudscheduler.admin`
- Cloudfunctions Developer: `roles/cloudfunctions.developer`
- IAM ServiceAccount User: `roles/iam.serviceAccountUser`

The [Project Factory module][project-factory-module] and the
[IAM module][iam-module] may be used in combination to provision a
service account with the necessary roles applied.

### APIs

A project with the following APIs enabled must be used to host the
resources of this module:

- Cloud Scheduler API: `cloudscheduler.googleapis.com`
- Cloud PubSub API: `pubsub.googleapis.com`
- Cloud Functions API: `cloudfunctions.googleapis.com`
- App Engine Admin API: `appengine.googleapis.com`

The [Project Factory module][project-factory-module] can be used to
provision a project with the necessary APIs enabled.

## Contributing

Refer to the [contribution guidelines](./CONTRIBUTING.md) for
information on contributing to this module.

[iam-module]: https://registry.terraform.io/modules/terraform-google-modules/iam/google
[project-factory-module]: https://registry.terraform.io/modules/terraform-google-modules/project-factory/google
[terraform-provider-gcp]: https://www.terraform.io/docs/providers/google/index.html
[terraform]: https://www.terraform.io/downloads.html
[cloud-scheduler-documentation]: https://cloud.google.com/scheduler/docs/
