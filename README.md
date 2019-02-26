# Terraform Network Module
This modules makes it easy to set up a scheduled job to trigger events/run functions.

## Usage
You can go to the examples folder, however the usage of the module could be like this in your own main.tf file:

```hcl
module "scheduled-function" {
  source  = "terraform-google-modules/scheduled-functions/google"
  version = "0.1.0"
  project_id   = "<PROJECT ID>"
  job_name="<NAME_OF_JOB>"
  schedule="<CRON_SYNTAX_SCHEDULE"
  function_entry_point="<NAME_OF_FUNCTION>"
  function_source_directory="<DIRECTORY_OF_FUNCTION_SOURCE>"
  name="<RESOURCE_NAMES>"
  region="<REGION>"
}
```

Then perform the following commands on the root folder:

- `terraform init` to get the plugins
- `terraform plan` to see the infrastructure plan
- `terraform apply` to apply the infrastructure build
- `terraform destroy` to destroy the built infrastructure

[^]: (autogen_docs_start)

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| function\_available\_memory\_mb | The amount of memory in megabytes allotted for the function to use. | string | `"256"` | no |
| function\_description | The description of the function. | string | `"Processes log export events provided through a Pub/Sub topic subscription."` | no |
| function\_entry\_point | The name of a method in the function source which will be invoked when the function is executed. | string | n/a | yes |
| function\_environment\_variables | A set of key/value environment variable pairs to assign to the function. | map | `<map>` | no |
| function\_event\_trigger\_failure\_policy\_retry | A toggle to determine if the function should be retried on failure. | string | `"false"` | no |
| function\_labels | A set of key/value label pairs to assign to the function. | map | `<map>` | no |
| function\_runtime | The runtime in which the function will be executed. | string | `"nodejs6"` | no |
| function\_source\_archive\_bucket\_labels | A set of key/value label pairs to assign to the function source archive bucket. | map | `<map>` | no |
| function\_source\_directory | The contents of this directory will be archived and used as the function source. | string | n/a | yes |
| function\_timeout\_s | The amount of time in seconds allotted for the execution of the function. | string | `"60"` | no |
| job\_description | Addition text to describet the job | string | `""` | no |
| job\_name | The name of the scheduled job to run | string | n/a | yes |
| job\_schedule | The job frequency, in cron syntax | string | `"*/2 * * * *"` | no |
| name | The name to apply to any nameable resources. | string | n/a | yes |
| project\_id | The ID of the project where this VPC will be created | string | n/a | yes |
| region | The region in which resources will be applied. | string | n/a | yes |
| topic\_name | Name of pubsub topic connecting the scheduled job and the function | string | `"test-topic"` | no |

## Outputs

| Name | Description |
|------|-------------|
| name | The name of the job created |

[^]: (autogen_docs_end)

## Requirements
### Terraform plugins
- [Terraform](https://www.terraform.io/downloads.html) 0.10.x
- [terraform-provider-google](https://github.com/terraform-providers/terraform-provider-google) plugin v1.12.0

### Configure a Service Account
In order to execute this module you must have a Service Account with permissions to create PubSub topics, create and deploy Cloud Functions, and create a Cloud Scheduler Job.

### Enable API's
In order to operate with the Service Account you must activate the following API on the project where the Service Account was created:

- Cloud Scheduler API - cloudscheduler.googleapis.com
- Cloud PubSub API - pubsub.googleapis.com
- Cloud Functions API - cloudfunctions.googleapis.com

## Install

### Terraform
Be sure you have the correct Terraform version (0.11.x), you can choose the binary here:
- https://releases.hashicorp.com/terraform/

## File structure
The project has the following folders and files:

- /: root folder
- /examples: examples for using this module
- /test: Folders with files for testing the module (see Testing section on this file)
- /main.tf: main file for this module, contains all the resources to create
- /variables.tf: all the variables for the module
- /output.tf: the outputs of the module
- /README.md: this file

## Testing and documentation generation

### Requirements
- [docker](https://docker.com)
- [terraform-docs](https://github.com/segmentio/terraform-docs/releases) 0.3.0

### Integration test
##### Terraform integration tests
It is recommended to to run the integration tests via docker. To do so, run `make test_integration_docker`. In containers, this will
- Perform `terraform init` command
- Perform `terraform get` command
- Perform `terraform plan` command and check that it'll create *n* resources, modify 0 resources and delete 0 resources
- Perform `terraform apply -auto-approve` command and check that it has created the *n* resources, modified 0 resources and deleted 0 resources
- Perform several `gcloud` commands and check the infrastructure is in the desired state
- Perform `terraform destroy -force` command and check that it has destroyed the *n* resources

### Autogeneration of documentation from .tf files
Run
```
make generate_docs
```

### Linting
The makefile in this project will lint or sometimes just format any shell,
Python, golang, Terraform, or Dockerfiles. The linters will only be run if
the makefile finds files with the appropriate file extension.

All of the linter checks are in the default make target, so you just have to
run

```
make -s
```

The -s is for 'silent'. Successful output looks like this

```
Running shellcheck
Running flake8
Running gofmt
Running terraform validate
Running hadolint on Dockerfiles
Test passed - Verified all file Apache 2 headers
```

The linters
are as follows:
* Shell - shellcheck. Can be found in homebrew
* Python - flake8. Can be installed with 'pip install flake8'
* Golang - gofmt. gofmt comes with the standard golang installation. golang
is a compiled language so there is no standard linter.
* Terraform - terraform has a built-in linter in the 'terraform validate'
command.
* Dockerfiles - hadolint. Can be found in homebrew
