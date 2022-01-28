package models
import(
"github.com/aws/aws-sdk-go/service/dynamodb"
"github.com/aws/aws-sdk-go/aws"
)
func TableModel()  (*dynamodb.CreateTableInput,string) {

	tableName := "Usuarios"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("email"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("email"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}


	return input, tableName


}