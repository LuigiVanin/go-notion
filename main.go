package main

import (
	"fmt"
	"log"
	"main/app"
	config "main/config"
	conn "main/config/database"
)



func main() {
	env, err := config.LoadEnvData()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	url := env.DatabaseURL

	db, err	:= conn.CreateConnection(url)

	fmt.Println("DATABASE: ", conn.Database)
	
	if err != nil {
		fmt.Println(db)
		log.Fatal("Error connecting to database")
		return
	}

	// migrationErr := config.Migrate(db)

	// if migrationErr != nil {
	// 	fmt.Println("Error migrating database")
	// }
	app.Run(&env)
}