package main

import (
	"fmt"
	"log"
	"main/app"
	config "main/config"
	conn "main/config/database"
)

// @title           Swagger Go Notion API
// @version         1.0
// @description     A simple swagger documentation for my go api.

// @contact.name   Luis Felipe Vanin
// @contact.email  luisfvaninmartins@gmail.com

// @host      localhost:3000
// @BasePath  /
func main() {
	env, err := config.LoadEnvData()
	fmt.Println("ENV: ", config.GlobalEnv)

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	url := env.DatabaseURL

	connection, err := conn.CreateConnection(url)

	fmt.Println("DATABASE: ", conn.Database)

	if err != nil {
		fmt.Println(connection)
		log.Fatal("Error connecting to database")
		return
	}

	migrationErr := conn.Migrate()

	if migrationErr != nil {
		fmt.Println("Error migrating database")
	}

	db, _ := connection.DB()
	defer db.Close()
	app.Run(&env)
}
