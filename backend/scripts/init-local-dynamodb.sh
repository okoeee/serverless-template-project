#!/usr/bin/env bash

REGION="ap-northeast-1"
ENDPOINT_URL="http://localhost:8000"
TABLE_NAME="dev-users"
TASKS_TABLE_NAME="dev-tasks"

# Create Users table
aws dynamodb create-table \
  --region "${REGION}" \
  --endpoint-url "${ENDPOINT_URL}" \
  --table-name "${TABLE_NAME}" \
  --attribute-definitions \
      AttributeName=userId,AttributeType=S \
      AttributeName=email,AttributeType=S \
  --key-schema \
      AttributeName=userId,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --global-secondary-indexes '[
    {
      "IndexName": "EmailIndex",
      "KeySchema": [
        { "AttributeName": "email", "KeyType": "HASH" }
      ],
      "Projection": { "ProjectionType": "ALL" }
    }
  ]'

# Create Task table 
aws dynamodb create-table \
  --region "${REGION}" \
  --endpoint-url "${ENDPOINT_URL}" \
  --table-name "${TASKS_TABLE_NAME}" \
  --attribute-definitions \
      AttributeName=taskId,AttributeType=S \
      AttributeName=userId,AttributeType=S \
  --key-schema \
      AttributeName=taskId,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --global-secondary-indexes '[
    {
      "IndexName": "UserTasksIndex",
      "KeySchema": [
        { "AttributeName": "userId", "KeyType": "HASH" }
      ],
      "Projection": { "ProjectionType": "ALL" }
    }
  ]'

echo "DynamoDB tables created successfully."
