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
| `BILLING_ACCOUNT` | Billing Account used to provision resources. | `string` | n/a | no |
| `BILLING_SINKS_PAGE_SIZE ` | The maximum number of Billing Account Log Sinks to return in the call to `BillingAccountsSinksService.List` service. | `number` | n/a | yes |
| `CLEAN_UP_BILLING_SINKS` | Clean up Billing Account Sinks. | `bool` | n/a | yes |
| `CLEAN_UP_CAI_FEEDS`| Clean up organization level Cloud Asset Inventory Feeds. | `bool` | n/a | yes |
| `CLEAN_UP_SCC_NOTIFICATIONS` | Clean up organization level Security Command Center notifications. | `bool` | n/a | yes |
| `CLEAN_UP_TAG_KEYS` | Clean up organization level Tag Keys. | `bool` | n/a | yes |
| `MAX_PROJECT_AGE_HOURS` | The project age, in hours, at which point deletion should be considered | integer | n/a | yes |
| `SCC_NOTIFICATIONS_PAGE_SIZE` | The maximum number of notification configs to return in the call to `ListNotificationConfigs` service. The minimun value is 1 and the maximum value is 1000. | `number` | n/a | yes |
| `TARGET_BILLING_SINKS` | List of Billing Account Log Sinks names regex that will be deleted. Regex example: `.*/sinks/sk-c-logging-.*-billing-.*` | `list(string)` | n/a | no |
| `TARGET_EXCLUDED_LABELS` | Labels to match on for identifying projects to avoid deletion | string | n/a | no |
| `TARGET_EXCLUDED_TAGKEYS` | List of organization Tag Key short names that won't be deleted. | `list(string)` | n/a | no |
| `TARGET_FOLDER_ID` | Folder ID to delete projects under | string | n/a | yes |
| `TARGET_INCLUDED_FEEDS` | List of organization level Cloud Asset Inventory feeds that should be deleted. Regex example: `.*/feeds/fd-cai-monitoring-.*` | `list(string)` | n/a | no |
| `TARGET_INCLUDED_LABELS` | Labels to match on for identifying projects to delete | string | n/a | no |
| `TARGET_INCLUDED_SCC_NOTIFICATIONS` | List of organization Security Command Center notifications names regex that will be deleted. Regex example: `.*/notificationConfigs/scc-notify-.*` | `list(string)` | n/a | no |
| `TARGET_ORGANIZATION_ID` | The organization ID whose projects to clean up | `string` | n/a | yes |

## Required Permissions

This Cloud Function must be run as a Service Account with the `Organization Administrator` (`roles/resourcemanager.organizationAdmin`) role.
If `CLEAN_UP_BILLING_SINKS` is enabled the Service Account running the Cloud Function needs role Logs Configuration Writer(`roles/logging.configWriter`) in the billing account `BILLING_ACCOUNT`.
