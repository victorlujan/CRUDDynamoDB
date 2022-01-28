package models
import(
"github.com/aws/aws-sdk-go/service/dynamodb"
"github.com/aws/aws-sdk-go/aws"
)
func TableModel()  (*dynamodb.CreateTableInput,string) {

	tableName := "Users"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Email"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Email"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(20),
			WriteCapacityUnits: aws.Int64(20),
		},
		TableName: aws.String(tableName),
	}


	return input, tableName


}