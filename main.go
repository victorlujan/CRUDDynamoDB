package main

import (
	"CRUDDynamoDB/CRUD"
	"CRUDDynamoDB/conexion"
	"CRUDDynamoDB/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	svc, _ := conexion.ConectDB()
	input, tableName := models.TableModel()

	_, err := CRUD.CreateTable(svc, input)

	if err == nil {
		fmt.Println("Created the table", tableName)

	} else {
		log.Fatal(err, " Error creating table")
	}
	
	CRUD.PutItem(svc, getItems(), tableName)

	//input := models.GetItemModel()
	//CRUD.GetItem(svc, input, tableName)

	//out, err :=CRUD.GetItem(svc, input)

	scaninput := models.GetScanItemModel()

	out, err := CRUD.ScanItem(svc, scaninput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}

func getItems() []models.User {
	raw, err := ioutil.ReadFile("models/user_data.json")
	if err != nil {
		log.Fatalf("Got error reading file: %s", err)
	}

	var items []models.User
	json.Unmarshal(raw, &items)
	return items
}
