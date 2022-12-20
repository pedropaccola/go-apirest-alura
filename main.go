package main

import (
	"fmt"

	"github.com/pedropaccola/go-restapi-alura/models"
	"github.com/pedropaccola/go-restapi-alura/routes"
	"github.com/pedropaccola/go-restapi-alura/database"
)

func main() {
	models.People = []models.Person{
		{ID: 1, Name: "Name01", History: "History01"},
		{ID: 2, Name: "Name02", History: "History02"},
	}

    database.Connect()

	fmt.Println("Starting REST server")
	routes.HandleRequest()
	fmt.Println("REST server started")
}
