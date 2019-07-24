variable "cron_schedule" {
  description = "Cron schedule for Google Cloud Scheduler"
  default = "55 * * * *"
}

variable "loader_alert_cloud_function_bucket"{
  description = "GCS bucket where python script + requirements.txt is located"
}

variable "loader_alert_cloud_function_zip"{
  description = "Zip file in Cloud Function Bucket which contains requirements.txt and python script"
  default = "audit_slack_alerts.zip"
}