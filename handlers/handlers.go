package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fransempe/backendGo/models"
	"github.com/fransempe/backendGo/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func Handlers() {
	getEntities()
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

type Entities struct {
	Entities []models.Entity `json:"entities"`
}

var JsonEntities Entities

func getEntities() {

	jsonFile, err := os.Open("db/data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database charged correctly.")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var entities Entities

	json.Unmarshal(byteValue, &entities)

	JsonEntities = entities

	/*for i := 0; i < len(entities.Entities); i++ {
		fmt.Println("----------------------------------")
		fmt.Println(entities.Entities[i].ID, "-", entities.Entities[i].Title)
		fmt.Println("Tipo de propiedad: ", entities.Entities[i].Property_Type.Name)
		fmt.Println("Moneda: ", entities.Entities[i].Currency.Name, "/ Estado: ", entities.Entities[i].Status)
	}*/
}
