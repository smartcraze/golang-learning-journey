package main

import (
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
func main() {
	app := gin.Default()

	app.GET("/", hello)
	app.Run(":8080")
}
