package main

import (
	"recipe-back/initializers"
	"recipe-back/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.SavedRecipe{})
}
