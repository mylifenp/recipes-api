package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mylifenp/recipes-api/api/v1/handlers"
)

func SetupTest(router *gin.Engine) {
	router.GET("/health", handlers.GetHealth)
}
