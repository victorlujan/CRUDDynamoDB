package CRUD

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteItem(svc *dynamodb.DynamoDB, input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	out, err := svc.DeleteItem(input)
	if err != nil {
		return nil, err
	}
	return out, nil
}
