package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getCount() uint64 {
	return 1337
}

func getTable() []string {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	client := dynamodb.NewFromConfig(cfg)

	tables, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println(err)
	}

	r := ""
	for _, name := range tables.TableNames {
		r = r + name
	}

	return tables.TableNames
}
