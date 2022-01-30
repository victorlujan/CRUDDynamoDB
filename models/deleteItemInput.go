package models

import(
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

func GetDeleteInput() *dynamodb.DeleteItemInput {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String("61f5b92b082bc2f68ecad05d"),
			},
			"Email": {
				S: aws.String("sophiamckenzie@comstruct.com"),
			},
		},
		TableName: aws.String("Users"),
	}
	return input
}