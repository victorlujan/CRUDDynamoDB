package models

import(
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

func GetItemModel()(*dynamodb.GetItemInput){

 tableName:="Users"	

input:= &dynamodb.GetItemInput{
	Key: map[string]*dynamodb.AttributeValue{
		"Id": {
			S: aws.String("2"),
		},
		"Email": {
			S: aws.String("example@gmail.com"),
		},
	},
	TableName: aws.String(tableName),
	
}


return input
}