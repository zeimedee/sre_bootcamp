package main

import (
	"os"

	"github.com/zeimedee/sre_bootcamp/internal/routes"
)

func main() {
	port := os.Getenv("PORT")

	router := routes.SetUpRouter()

	router.Run(port)
}
