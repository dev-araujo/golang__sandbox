package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Nossa API de TODOs com Gin está no ar!",
		})
	})

	router.Run(":8080")
}
