package main

import (
	"github.com/mylifenp/recipes-api/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run("localhost:8080")
}
