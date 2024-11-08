package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zeimedee/sre_bootcamp/internal/database"
	"github.com/zeimedee/sre_bootcamp/internal/routes"
)

func main() {
	er := godotenv.Load()
	if er != nil {
		log.Fatalf("Error loading .env file: %v", er)
	}
	port := os.Getenv("PORT")
	database.ConnectDb()
	router := routes.SetUpRouter()

	router.Run(port)
}
