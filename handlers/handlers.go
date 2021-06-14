package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fransempe/backendGo/authentication"
	"github.com/fransempe/backendGo/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/login", authentication.Login).Methods("POST")
	router.HandleFunc("/entities", routes.GetProperties).Methods("GET")
	router.HandleFunc("/entitiess", routes.GetPropertyById).Methods("GET")
	router.HandleFunc("/entities2", routes.NewProperty).Methods("POST")
	router.HandleFunc("/entities3", routes.DeleteProperty).Methods("POST")
	router.HandleFunc("/entities4", routes.ChangePropertyStatus).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	corsAllow := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, corsAllow))
}
