package main

import (
	"api-rest-go/models"
	"api-rest-go/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "John", Historia: "Grande cantor"},
		{Id: 2, Nome: "Josue", Historia: "Grande construtor"},
	}

	routes.HandleRequest()
}
