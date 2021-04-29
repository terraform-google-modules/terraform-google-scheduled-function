# Terraform Google Scheduled Functions Module

This modules makes it easy to set up a scheduled job to trigger events/run functions.

## Compatibility
This module is meant for use with Terraform 0.13. If you haven't
[upgraded](https://www.terraform.io/upgrade-guides/0-13.html) and need a Terraform
0.12.x-compatible version of this module, the last released version
intended for Terraform 0.12.x is [v1.5.1](https://registry.terraform.io/modules/terraform-google-modules/-scheduled-function/google/v1.5.1).

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
