package main

import (
	"fmt"
	"net/http"

	"day5-http-server/routes"
)

func main() {
	routes.RegisterUserRoutes()

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
