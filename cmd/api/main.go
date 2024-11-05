package main

import (
	"os"

	"github.com/zeimedee/sre_bootcamp/internal/database"
	"github.com/zeimedee/sre_bootcamp/internal/routes"
)

func main() {

	port := os.Getenv("PORT")
	database.ConnectDb()
	router := routes.SetUpRouter()

	router.Run(port)
}
