package conexion

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Conectar() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "eu-west-1"
		return nil
	})

	return dynamodb.NewFromConfig(cfg), err
}
