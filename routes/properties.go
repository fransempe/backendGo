package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fransempe/backendGo/authentication"
	"github.com/fransempe/backendGo/models"
)

//var PropertiesJson models.Property = ParsePropertiesJson()

func GetProperties(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)

}

func GetPropertyById(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)
	/*vars := mux.Vars(r)
	propertyID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, property := range properties {
		if property.ID == propertyID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(property)
		}
	}*/

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

	/*vars := mux.Vars(r)
	propertyID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, property := range properties {
		if property.ID == propertyID {

			properties = append(properties[:i], properties[i+1:]...)
			fmt.Fprintf(w, "The property with ID %v hast been removed succesfully")
		}
	}*/
}

func ChangePropertyStatus(w http.ResponseWriter, r *http.Request) {
	authentication.VerifyToken(w, r)

	/*vars := mux.Vars(r)
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
		json.Unmarshal(reqBody, &updatedProperty)

		for i, property := range properties {
			if property.ID == propertyID {

				properties = append(properties[:i], properties[i+1:]...)
				updatedProperty.Status
			}
		}*/
}
