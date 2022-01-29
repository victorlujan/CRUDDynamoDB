package CRUD

import (
	"CRUDDynamoDB/models"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetItem(svc *dynamodb.DynamoDB, input *dynamodb.GetItemInput) (models.User, error) {
	result, err := svc.GetItem(input)
	if err != nil {
		
		return models.User{}, err
	}

	item:= models.User{}


	err = dynamodbattribute.UnmarshalMap(result.Item,&item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}






	return item , nil

}
