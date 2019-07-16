package main

import (
	"io"
	"log"
	"net/http"

	"github.com/flume-cloud-services/cache/middleware"
	"github.com/flume-cloud-services/database/controllers"
)

func main() {
	http.HandleFunc("/signin", controllers.Signin)

	http.Handle("/welcome", middleware.Middleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, "Welcome to you visitor !")
		}),
		middleware.AuthMiddleware,
	))

	log.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
