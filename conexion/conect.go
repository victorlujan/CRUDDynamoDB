package conexion

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

func ConectDB() (*dynamodb.DynamoDB, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("eu-west-1"),
		},
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	err := errors.New("Error al conectar con la base de datos")

	return svc , err

}
