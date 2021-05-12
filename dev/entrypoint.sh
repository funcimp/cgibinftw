#!/bin/bash

PATH=/usr/local/aws-cli/v2/2.2.3/bin:$PATH
TABLE_NAME="ulticntr"

aws dynamodb create-table \
    --endpoint-url "${ENDPOINT_URL}" \
    --table-name "${TABLE_NAME}" \
    --key-schema AttributeName=counter_id,KeyType=HASH \
    --attribute-definitions AttributeName=counter_id,AttributeType=S \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

aws dynamodb put-item \
    --endpoint-url "${ENDPOINT_URL}" \
    --table-name "${TABLE_NAME}" \
    --item '{ "counter_id": {"S": "primary"}, "hit": {"N": "0"} }'

while true
do
	sleep 1
done