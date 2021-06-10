package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fransempe/backendGo/handlers"
	"github.com/fransempe/backendGo/models"
)

type Entities struct {
	Entities []models.Entity `json:"entities"`
}

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

	/*for i := 0; i < len(entities.Entities); i++ {
		fmt.Println("----------------------------------")
		fmt.Println(entities.Entities[i].Id, "-", entities.Entities[i].Title)
		fmt.Println("Tipo de propiedad: ", entities.Entities[i].Property_Type.Name)
		fmt.Println("Moneda: ", entities.Entities[i].Currency.Name, "/ Estado: ", entities.Entities[i].Status)
	}*/

}

func main() {
	getEntities()
	handlers.Handlers()
}
