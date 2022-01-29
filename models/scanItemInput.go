package models
import(
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"

)

func GetScanItemModel()(*dynamodb.ScanInput){
	
	tableName:="Users"	
	
	input:= &dynamodb.ScanInput{
		
		TableName: aws.String(tableName),
		
	}
	
	
	return input
	
	
}