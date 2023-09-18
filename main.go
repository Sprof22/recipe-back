package main

import (
	"recipe-back/controllers"
	"recipe-back/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Add your frontend URL
	r.Use(cors.New(config))

	r.POST("/recipes", controllers.CreateRecipe)
	r.GET("/recipes", controllers.GetRecipe)
	r.GET("/recipes/:id", controllers.GetSingleRecipe)
	r.PUT("/recipes/:id", controllers.UpdateRecipe)
	r.DELETE("/recipes/:id", controllers.DeleteRecipe)
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080
}
