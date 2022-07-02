package main

import (
	"fmt"

	"github.com/jheniferms/alura_api_rest/models"
	"github.com/jheniferms/alura_api_rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{"nome1", "historia 1"},
		{"nome2", "historia 2"},
	}

	fmt.Println("Iniciando servidor rest com go")

	routes.HandleRequest()
}