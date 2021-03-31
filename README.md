# Terraform Google Scheduled Functions Module

This modules makes it easy to set up a scheduled job to trigger events/run functions.

## Compatibility

This module is meant for use with Terraform 0.12. If you haven't
[upgraded](https://www.terraform.io/upgrade-guides/0-12.html) and need a Terraform 0.11.x-compatible
version of this module, the last released version intended for Terraform 0.11.x
is [v0.4.1](https://registry.terraform.io/modules/terraform-google-modules/scheduled-function/google/0.4.1).

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
|------|-------------|:----:|:-----:|:-----:|
| bucket\_force\_destroy | When deleting the GCS bucket containing the cloud function, delete all objects in the bucket first. | bool | `"true"` | no |
| bucket\_name | The name to apply to the bucket. Will default to a string of <project-id>-scheduled-function-XXXX> with XXXX being random characters. | string | `""` | no |
| function\_available\_memory\_mb | The amount of memory in megabytes allotted for the function to use. | number | `"256"` | no |
| function\_description | The description of the function. | string | `"Processes log export events provided through a Pub/Sub topic subscription."` | no |
| function\_entry\_point | The name of a method in the function source which will be invoked when the function is executed. | string | n/a | yes |
| function\_environment\_variables | A set of key/value environment variable pairs to assign to the function. | map(string) | `<map>` | no |
| function\_event\_trigger\_failure\_policy\_retry | A toggle to determine if the function should be retried on failure. | bool | `"false"` | no |
| function\_labels | A set of key/value label pairs to assign to the function. | map(string) | `<map>` | no |
| function\_name | The name to apply to the function | string | n/a | yes |
| function\_runtime | The runtime in which the function will be executed. | string | `"nodejs10"` | no |
| function\_service\_account\_email | The service account to run the function as. | string | `""` | no |
| function\_source\_archive\_bucket\_labels | A set of key/value label pairs to assign to the function source archive bucket. | map(string) | `<map>` | no |
| function\_source\_dependent\_files | A list of any terraform created `local_file`s that the module will wait for before creating the archive. | object | `<list>` | no |
| function\_source\_directory | The contents of this directory will be archived and used as the function source. | string | n/a | yes |
| function\_timeout\_s | The amount of time in seconds allotted for the execution of the function. | number | `"60"` | no |
| grant\_token\_creator | Specify true if you want to add token creator role to the default Pub/Sub SA | bool | `"false"` | no |
| job\_description | Addition text to describe the job | string | `""` | no |
| job\_name | The name of the scheduled job to run | string | `"null"` | no |
| job\_schedule | The job frequency, in cron syntax | string | `"*/2 * * * *"` | no |
| message\_data | The data to send in the topic message. | string | `"dGVzdA=="` | no |
| project\_id | The ID of the project where the resources will be created | string | n/a | yes |
| region | The region in which resources will be applied. | string | n/a | yes |
| scheduler\_job | An existing Cloud Scheduler job instance | object | `"null"` | no |
| time\_zone | The timezone to use in scheduler | string | `"Etc/UTC"` | no |
| topic\_name | Name of pubsub topic connecting the scheduled job and the function | string | `"test-topic"` | no |

## Outputs

| Name | Description |
|------|-------------|
| name | The name of the job created |
| pubsub\_topic\_name | PubSub topic name |
| scheduler\_job | The Cloud Scheduler job instance |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

These sections describe requirements for using this module.

### Software

The following dependencies must be available:

- [Terraform][terraform] v0.12
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
