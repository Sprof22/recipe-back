package controllers

import (
	"recipe-back/initializers"
	"recipe-back/models"

	"github.com/gin-gonic/gin"
)

func CreateRecipe(c *gin.Context) {

	var body struct {
		Title        string
		Image        string
		Instructions string
	}

	c.Bind(&body)
	recipe := models.SavedRecipe{Title: body.Title, Image: body.Image, Instructions: body.Instructions}

	result := initializers.DB.Create(&recipe)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"recipe": recipe,
	})
}

func GetSingleRecipe(c *gin.Context) {

	id := c.Param("id")
	var recipe models.SavedRecipe
	initializers.DB.First(&recipe, id)

	c.JSON(200, gin.H{
		"recipe": recipe,
	})
}

func GetRecipe(c *gin.Context) {
	var recipes []models.SavedRecipe
	result := initializers.DB.Find(&recipes)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"recipes": recipes,
	})
}

func UpdateRecipe(c *gin.Context) {
	var body struct {
		Title        string
		Instructions string
		Image        string
	}

	id := c.Param("id")
	var recipe models.SavedRecipe
	initializers.DB.First(&recipe, id)
	c.Bind(&body)

	initializers.DB.Model(&recipe).Updates(models.SavedRecipe{Title: body.Title, Instructions: body.Instructions, Image: body.Image})

	c.JSON(200, gin.H{

		"recipe": recipe,
	})

}

func DeleteRecipe(c *gin.Context) {
	id := c.Param("id")
	var actor models.SavedRecipe
	initializers.DB.Delete(&models.SavedRecipe{}, id)

	c.JSON(200, gin.H{
		"actor": actor,
	})
}
