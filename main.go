package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/swaggo/gin-swagger/example/multiple/api/v1"
)

// @title Recipes Example
// @version 1.0
// @description This is a sample server Recipe server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	router := gin.New()
	v1.SetupRecipe(router)
	v1.SetupTest(router)
	// router := routes.SetupRouter()
	// docs.SwaggerInfo.BasePath = "/"
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/docs/v1/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080")
}
