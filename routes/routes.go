package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mylifenp/recipes-api/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/recipes", handlers.ListRecipes)
	return router
}
