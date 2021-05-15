package counter

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	tableName  = "ulticntr"
	counterCol = "counter_id"
	counterRow = "primary"
)

type dynamo struct {
	client *dynamodb.Client
}

// Count is dynamo's implementation of the Counter interface.
func (d dynamo) Count() (v uint64, err error) {
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

func newDynamo() (*dynamo, error) {
	client, err := newClient()
	return &dynamo{client: client}, err
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
	t := tableName
	updateExpression := "SET hit = hit + :incr"
	key := make(map[string]types.AttributeValue)
	key[counterCol] = &types.AttributeValueMemberS{Value: counterRow}
	expAttVals := make(map[string]types.AttributeValue)
	expAttVals[":incr"] = &types.AttributeValueMemberN{Value: "1"}

	return &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 &t,
		UpdateExpression:          &updateExpression,
		ExpressionAttributeValues: expAttVals,
		ReturnValues:              "UPDATED_NEW",
	}
}
