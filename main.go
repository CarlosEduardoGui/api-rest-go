package main

import (
	"api-rest-go/database"
	"api-rest-go/models"
	"api-rest-go/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "John", Historia: "Grande cantor"},
		{Id: 2, Nome: "Josue", Historia: "Grande construtor"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
