package main

import (
	"fmt"
	"go-boilerplate/app"
	"go-boilerplate/db"
	"go-boilerplate/router"
)

func main() {
	// Carrega dados de configuração
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// Conecta ao banco de dados
	database := db.Connect()
	defer database.Close()

	// Roda o servidor
	routers := router.Setup(database)
	routers.Run(fmt.Sprintf(":%v", app.Config.ServerPort))
}
