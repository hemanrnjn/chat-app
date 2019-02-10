package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hemanrnjn/chat-app/controllers"

	"github.com/hemanrnjn/chat-app/app"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

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

	err := http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
