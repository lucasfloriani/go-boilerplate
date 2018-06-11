package main

import (
	"fmt"
	"go-boilerplate/app"
	"go-boilerplate/db"
	"go-boilerplate/router"
)

func main() {
	// Loads configuration data
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// Connects to the database
	database := db.Connect()
	defer database.Close()

	// Runs the server
	routers := router.Setup(database)
	routers.Run(fmt.Sprintf(":%v", app.Config.ServerPort))
}
