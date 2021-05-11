package main

import (
	"context"
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func newClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	endpoint := dynamodb.EndpointResolverFromURL("http://dynamodb-local:8000")
	return dynamodb.NewFromConfig(cfg, dynamodb.WithEndpointResolver(endpoint)), err
}

func updateCounterInput() *dynamodb.UpdateItemInput {
	tableName := "ulticntr"
	updateExpression := "SET hit = hit + :incr"
	key := make(map[string]types.AttributeValue)
	key["counter_id"] = &types.AttributeValueMemberS{Value: "primary"}
	expAttVals := make(map[string]types.AttributeValue)
	expAttVals[":incr"] = &types.AttributeValueMemberN{Value: "1"}

	return &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 &tableName,
		UpdateExpression:          &updateExpression,
		ExpressionAttributeValues: expAttVals,
		ReturnValues:              "UPDATED_NEW",
	}
}

func logVisit() (v uint64, err error) {
	ctx := context.Background()
	client, err := newClient(ctx)
	if err != nil {
		return v, err
	}
	result, err := client.UpdateItem(ctx, updateCounterInput())
	if err != nil {
		return v, err
	}
	output, ok := result.Attributes["hit"].(*types.AttributeValueMemberN)
	if !ok {
		return v, errors.New("typecasting failed for hit")
	}
	return strconv.ParseUint(output.Value, 10, 64)
}
