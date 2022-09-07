package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mylifenp/recipes-api/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/recipes", handlers.NewRecipe)
	router.GET("/recipes", handlers.ListRecipes)
	router.PUT("/recipes/:id", handlers.UpdateRecipe)
	router.DELETE("/recipes/:id", handlers.DeleteRecipe)
	router.GET("/recipes/search", handlers.SearchRecipes)
	return router
}
