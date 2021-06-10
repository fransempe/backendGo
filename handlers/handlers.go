package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	router.HandleFunc("/login", routes.Login).Methods("POST")
	router.HandleFunc("/entities", routes.EntitiesList).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	corsAllow := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, corsAllow))
}
