
resource "google_cloud_scheduler_job" "cron-audit-logs-check"{
	name = "audit-logs-error-check"
	region = "us-central1"
	description = "Scheduled time to run audit query to check for errors"
	schedule = "${var.cron_schedule}"
	http_target = {
    		http_method = "POST"
    		uri = "${google_cloudfunctions_function.audit-log.https_trigger_url}"
  }


}

resource "google_cloudfunctions_function" "audit-log" {

	name = "audit-log-slack-alert"
	region = "us-central1"
	description = "Cloud Function to query audit logs for loader errors"
	source_archive_bucket = "${var.loader_alert_cloud_function_bucket}" 
	source_archive_object = "${var.loader_alert_cloud_function_zip}"
	runtime = "python37"
	entry_point = "query_for_errors"
	trigger_http = true

}