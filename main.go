package main

import (
	"CRUDDynamoDB/CRUD"
	"CRUDDynamoDB/conexion"
	"CRUDDynamoDB/models"
	"fmt"
)

func main() {
	svc, _ := conexion.ConectDB()
	input, tableName := models.TableModel()

	out, _ := CRUD.CreateTable(svc, input)

	if out != nil {
		fmt.Println("Created the table", tableName)
		
	}else {
		fmt.Println("Error")
	}

}
