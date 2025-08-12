package main

import (
	"todo-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.TodoRoutes(app)

	app.Run(":8080")
}
