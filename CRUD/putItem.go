package CRUD

import (
	"CRUDDynamoDB/models"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func PutItem(svc *dynamodb.DynamoDB, data []models.User, tableName string)  {

	for _, item := range data {
		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			fmt.Println("Got error marshalling map:")
			log.Fatal(err.Error())
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("Users"),
		}
		fmt.Println(input)
		out, err := svc.PutItem(input)

		if err != nil {
			fmt.Println("Got error calling PutItem:")
			log.Fatal(err.Error())
		}else{
			fmt.Println("PutItem", out)
		}
	}
}
