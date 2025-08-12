package routes

import (
	"net/http"

	"day5-http-server/handlers"
)

func RegisterUserRoutes() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetUsers(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateUser(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
