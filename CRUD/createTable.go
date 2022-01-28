package CRUD
import(
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

func CreateTable(svc *dynamodb.DynamoDB, input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)   {

	out,err := svc.CreateTable(input)

	if err != nil {
		return nil, err
	}

	return out,nil
}