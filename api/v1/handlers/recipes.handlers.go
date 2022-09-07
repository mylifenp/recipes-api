package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mylifenp/recipes-api/api/v1/models"
	"github.com/rs/xid"
)

var recipes []models.Recipe

func init() {
	recipes = make([]models.Recipe, 0)
	file, _ := ioutil.ReadFile("./recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

// @Summary list all recipes in the recipes list
// @ID list-all-recipes
// @Produce json
// @Success 200 {object} models.Recipe
// @Router /recipes [get]
func ListRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func NewRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func UpdateRecipe(c *gin.Context) {
	id := c.Param("id")
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)
}

func DeleteRecipe(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been deleted"})
}

func SearchRecipes(c *gin.Context) {
	tag := c.Query("tag")
	listOfRecipes := make([]models.Recipe, 0)

	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes,
				recipes[i])
		}
	}
	c.JSON(http.StatusOK, listOfRecipes)
}
