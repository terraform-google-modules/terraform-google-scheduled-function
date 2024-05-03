# Project Cleanup Utility

This is a simple utility that scans a GCP organization for projects matching certain criteria, and enqueues such projects for deletion. Currently supported criteria are the combination of:

- **Age:** Only projects older than the configured age, in hours, will be marked for deletion.
- **Key-Value Pair Include:** Only projects whose labels contain the provided key-value pair will be marked for deletion.
- **Key-Value Pair Exclude:** Projects whose labels contain the provided key-value pair won't be marked for deletion.
- **Folder ID:** Only projects under this Folder ID will be recursively marked for deletion.

Both of these criteria must be met for a project to be deleted.

## Environment Configuration

The following environment variables may be specified to configure the cleanup utility:

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| `TARGET_EXCLUDED_LABELS` | Labels to match on for identifying projects to avoid deletion | string | n/a | no |
| `TARGET_FOLDER_ID` | Folder ID to delete projects under | string | n/a | yes |
| `TARGET_INCLUDED_LABELS` | Labels to match on for identifying projects to delete | string | n/a | no |
| `MAX_PROJECT_AGE_HOURS` | The project age, in hours, at which point deletion should be considered | integer | n/a | yes |

## Required Permissions

This Cloud Function must be run as a Service Account with the `Organization Administrator` (`roles/resourcemanager.organizationAdmin`) role.
