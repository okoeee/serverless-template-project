#!/usr/bin/env bash

REGION="ap-northeast-1"
ENDPOINT_URL="http://localhost:8000"
TABLE_NAME="dev-users"

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
