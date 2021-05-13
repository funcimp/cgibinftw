package main

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type dynamoCounter struct {
	client *dynamodb.Client
}

func newDynamoCounter() (*dynamoCounter, error) {
	client, err := newClient()
	return &dynamoCounter{client: client}, err
}

func newClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	var opts []func(*dynamodb.Options)

	if u := os.Getenv("ENDPOINT_URL"); u != "" {
		endpoint := dynamodb.EndpointResolverFromURL(u)
		o := dynamodb.WithEndpointResolver(endpoint)
		opts = append(opts, o)
	}

	return dynamodb.NewFromConfig(cfg, opts...), err
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

func (d dynamoCounter) Count() (v uint64, err error) {
	result, err := d.client.UpdateItem(context.Background(), updateCounterInput())
	if err != nil {
		return v, err
	}
	output, ok := result.Attributes["hit"].(*types.AttributeValueMemberN)
	if !ok {
		return v, errors.New("typecasting failed for hit")
	}
	return strconv.ParseUint(output.Value, 10, 64)
}
