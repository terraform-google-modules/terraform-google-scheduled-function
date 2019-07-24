import logging
import requests
from google.cloud import bigquery


logging.getLogger().setLevel(logging.INFO)
SLACK_WEBHOOK_URL = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXX"
QUERY = """
WITH
  errors AS (
  SELECT
    protopayload_auditlog.servicedata_v1_bigquery.jobInsertResponse.resource.jobStatus.error.message AS error_message,
    EXTRACT(HOUR FROM current_timestamp) as hr,
  FROM
    dataset.cloudaudit_googleapis_com_data_access
  WHERE
    protopayload_auditlog.servicedata_v1_bigquery.jobInsertResponse.resource.jobStatus.error.message IS NOT NULL
    AND EXTRACT(HOUR
    FROM
      current_timestamp) = EXTRACT(HOUR
    FROM
      protopayload_auditlog.servicedata_v1_bigquery.jobInsertResponse.resource.jobStatistics.createTime))
SELECT
  error_message as Error,
  hr,
  COUNT(*) as Count
FROM
  errors
GROUP BY
  1,2

"""

def query_for_errors(incoming_request):

  bq_client = bigquery.Client()
  logging.info("Running: {0}".format(QUERY))
  query_job = bq_client.query(query)

  if list(query_job):
    for row in list(query_job):
      text = ("Alert: Error {0} from the {1} component has occurred {2} times"
      "in the past hour - {3}:00 PST. "
      "Please file a bug ticket to have)".format(
      str(row['Error'][:500]),
      str(row['Count']),
      str(row['hr']))
      logging.info("Posting to slack: {0}".format(text))
      r=requests.post(url=SLACK_URL_HOOK, data = str({"text": text}))
      logging.info(r.text)

if __name__ == "__main__":
    query_for_errors(None)
