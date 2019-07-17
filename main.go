package main

import (
	"io"
	"log"
	"net/http"

	"github.com/flume-cloud-services/cache/controllers"
	"github.com/flume-cloud-services/cache/middleware"
)

func main() {
	http.HandleFunc("/signin", controllers.Signin)

	http.Handle("/welcome", middleware.Middleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, "Welcome to you visitor !")
		}),
		middleware.AuthMiddleware,
	))

	http.Handle("/insert", middleware.Middleware(
		http.HandlerFunc(controllers.InsertData),
		middleware.AuthMiddleware),
	)

	log.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
