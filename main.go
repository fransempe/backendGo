package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fransempe/backendGo/authentication"
	"github.com/fransempe/backendGo/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*TODO: El proyecto esta separado por paquetes (como sugiere la documentación de GO)
en este caso, queda embebido dentro del package MAIN a fin de poder usar la variable JsonProperties (Es la variable que contendrá
el json provisto al iniciar la aplicación). De todas formas, presento la estructura del proyecto como debiera estar distribuida */

//Inicializo la variable con los datos provistos del json.
var JsonProperties []models.Property = ParsePropertiesJson()

func main() {
	Handlers()
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome GO|API")
}

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/login", authentication.Login).Methods("POST")
	router.HandleFunc("/properties", GetProperties).Methods("GET")
	router.HandleFunc("/properties/{id}", GetPropertyById).Methods("GET")
	router.HandleFunc("/properties", NewProperty).Methods("POST")
	router.HandleFunc("/properties/{id}", DeleteProperty).Methods("DELETE")
	router.HandleFunc("/properties/{id}", ChangePropertyStatus).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	corsAllow := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, corsAllow))
}

func GetProperties(w http.ResponseWriter, r *http.Request) {
	var verify bool = authentication.VerifyToken(w, r)
	if !verify {
		fmt.Fprintf(w, "Unauthorized")
	} else {
		decoder := json.NewDecoder(r.Body)
		var params string
		decoder.Decode(&params)
		if params != "" {
			for _, prop := range JsonProperties {
				if prop.Property_Type.Name == params {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(prop)
				}
			}
		} else {
			json.NewEncoder(w).Encode(JsonProperties)
		}

	}
}

func GetPropertyById(w http.ResponseWriter, r *http.Request) {
	var verify bool = authentication.VerifyToken(w, r)
	if !verify {
		fmt.Fprintf(w, "Unauthorized")
	} else {
		vars := mux.Vars(r)
		propertyID, err := strconv.Atoi(vars["id"])

		if err != nil {
			fmt.Fprintf(w, "Invalid ID")
			return
		}

		for _, prop := range JsonProperties {
			if prop.ID == propertyID {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(prop)
			}
		}
	}
}

func NewProperty(w http.ResponseWriter, r *http.Request) {
	var verify bool = authentication.VerifyToken(w, r)
	if !verify {
		fmt.Fprintf(w, "Unauthorized")
	} else {
		var newProperty models.Property

		reqBody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Fprintf(w, "Insert a new property")
		}

		json.Unmarshal(reqBody, &newProperty)

		//Autogenerar un id a partir de mi lista existente
		newProperty.ID = len(JsonProperties) + 1
		JsonProperties = append(JsonProperties, newProperty)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newProperty)
	}
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	var verify bool = authentication.VerifyToken(w, r)
	if !verify {
		fmt.Fprintf(w, "Unauthorized")
	} else {
		vars := mux.Vars(r)
		propertyID, err := strconv.Atoi(vars["id"])

		if err != nil {
			fmt.Fprintf(w, "Invalid ID")
			return
		}

		for i, property := range JsonProperties {
			if property.ID == propertyID {
				JsonProperties = append(JsonProperties[:i], JsonProperties[i+1:]...)
				fmt.Fprintf(w, "The property has been removed succesfully")
			}
		}
	}
}

func ChangePropertyStatus(w http.ResponseWriter, r *http.Request) {
	var verify bool = authentication.VerifyToken(w, r)
	if !verify {
		fmt.Fprintf(w, "Unauthorized")
	} else {
		decoder := json.NewDecoder(r.Body)

		// Se usa para almacenar datos de clave de parámetro = valor
		var params string // ingresar en el body el string: "available","rented" o "closed"

		// Analiza los parámetros en el mapa
		decoder.Decode(&params)

		vars := mux.Vars(r)
		propertyID, err := strconv.Atoi(vars["id"])

		if err != nil {
			fmt.Fprintf(w, "Invalid ID")
			return
		}

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Please enter valid data")
			return
		}
		json.Unmarshal(reqBody, &JsonProperties)

		for i, property := range JsonProperties {
			if property.ID == propertyID {
				JsonProperties[i].Status = params
			}
		}
	}
}

type Properties struct {
	Properties []models.Property `json:"properties"`
}

func ParsePropertiesJson() []models.Property {

	file, _ := ioutil.ReadFile("db/data.json")

	data := Properties{}

	_ = json.Unmarshal([]byte(file), &data)

	return data.Properties
}
