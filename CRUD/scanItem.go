package CRUD

import (
	"CRUDDynamoDB/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func ScanItem(svc *dynamodb.DynamoDB, input *dynamodb.ScanInput) ([]models.User, error) {
	result, err := svc.Scan(input)
	if err != nil {
		return []models.User{}, err
	}
	item := []models.User{}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &item)

	return item, nil
}
