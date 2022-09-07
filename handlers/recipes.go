package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mylifenp/recipes-api/models"
)

var recipes []models.Recipe

func init() {
	recipes = make([]models.Recipe, 0)
	file, _ := ioutil.ReadFile("../recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func ListRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}
