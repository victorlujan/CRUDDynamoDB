package main

import (
	"context"
	"log"
	"fmt"

	"CRUDDynamoDB/conexion"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodbstreams/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {

	dynamo, err := conexion.Conectar()

	if err != nil {
		panic(err)
	}

	proj := expression.NamesList(expression.Name("id"), expression.Name("email"), expression.Name("name"))



	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("my-table"),
	}

	type Item struct {
		id   int
		name  string
		email   string
	}






	out, err := dynamo.Scan(context.TODO(), params)

	if err != nil {
		panic(err)
	}

	numItems:= 0


	for _, i := range out.Items {
		item := Item{}
	
		err = attributevalue.UnmarshalMap(i, &item)
	
		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}
	
		// Which ones had a higher rating than minimum?
		
			// Or it we had filtered by rating previously:
			//   if item.Year == year {
			numItems++
	
			fmt.Println("Title: ", item.id)
			fmt.Println("Rating:", item.name)
			fmt.Println()
	}


}
