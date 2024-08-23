package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"main/app"
	"main/database/repository"

	_ "embed"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	err = repository.ConnectDB()
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}
	defer repository.DB.Close()

	fiberApp := fiber.New()

	app.SetupRoutes(fiberApp)

	err = fiberApp.Listen(":8080")
	if err != nil {
		return
	}

}
