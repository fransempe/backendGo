package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fransempe/backendGo/authentication"
	"github.com/fransempe/backendGo/models"
	"github.com/gorilla/mux"
)

var JsonProperties models.Properties

func GetProperties(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
	json.NewEncoder(w).Encode(JsonProperties)
}

func GetPropertyById(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
	vars := mux.Vars(r)
	propertyID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, prop := range JsonProperties.Properties {
		if prop.ID == propertyID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(prop)
		}
	}
}

func NewProperty(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
	var newProperty models.Property

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a new property")
	}

	json.Unmarshal(reqBody, &newProperty)

	//Autogenerar un id a partir de mi lista existente
	newProperty.ID = len(JsonProperties.Properties) + 1

	output := make([]models.Property, len(JsonProperties.Properties))

	output = append(output, newProperty)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
	/*vars := mux.Vars(r)
	propertyID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, property := range JsonProperties.Properties {
		if property.ID == propertyID {
			JsonProperties = append(JsonProperties.Properties[:i], JsonProperties.Properties[i+1:]...)
			fmt.Fprintf(w, "The property has been removed succesfully")
		}
	}*/
}

func ChangePropertyStatus(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
	decoder := json.NewDecoder(r.Body)

	// Se usa para almacenar datos de clave de parámetro = valor
	var params string

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

	for i, property := range JsonProperties.Properties {
		if property.ID == propertyID {
			JsonProperties.Properties[i].Status = params
		}
	}
}
