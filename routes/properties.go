package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../main"
	"github.com/fransempe/backendGo/authentication"
	"github.com/fransempe/backendGo/models"
)

//var PropertiesJson models.Property = ParsePropertiesJson()

func GetProperties(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
	fmt.Println(main.JsonStr)
	//ParsePropertiesJson()
}

func GetPropertyById(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
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
	//newProperty.ID = len(PropertiesJson) + 1
	//append(PropertiesJson, newProperty)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProperty)
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
}

func ChangePropertyStatus(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
}
