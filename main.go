package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hemanrnjn/chat-app/app"
	"github.com/hemanrnjn/chat-app/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/ws", controllers.HandleSocketMessages)
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	go controllers.HandleMessages()

	err := http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router)) //Launch the app, visit localhost:8000/api
	// err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
