import logging
import requests
from google.cloud import bigquery


logging.getLogger().setLevel(logging.INFO)

#Fill in variables with your project setup
variables ={
  "SLACK_WEBHOOK_URL": "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXX",
  "DATASET_NAME": "mydataset",
  "AUDIT_LOG_TABLE":  "audit_table_name",
  "TIME_COLUMN": "time_column",
  "ERROR_MESSAGE_COLUMN":"error_message_column"
}
QUERY = """
WITH
  errors AS (
  SELECT
    {ERROR_MESSAGE_COLUMN} AS error_message,
    EXTRACT(HOUR FROM current_timestamp) as hr,
  FROM
    {DATASET_NAME}.{AUDIT_LOG_TABLE}
  WHERE
    {ERROR_MESSAGE_COLUMN} IS NOT NULL
    AND EXTRACT(HOUR
    FROM
      current_timestamp) = EXTRACT(HOUR
    FROM
     {TIME_COLUMN}))
SELECT
  error_message as Error,
  hr,
  COUNT(*) as Count
FROM
  errors
GROUP BY
  1,2
""".format(**variables)

def query_for_errors(incoming_request):

  bq_client = bigquery.Client()
  logging.info("Running: {0}".format(QUERY))
  query_job = bq_client.query(QUERY)

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
