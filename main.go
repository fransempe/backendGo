package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fransempe/backendGo/handlers"
	"github.com/fransempe/backendGo/models"
)

func main() {
	var JsonStr string = ParsePropertiesJson()
	handlers.Handlers()
}

//Acá parseo el Json y lo paso a una variable global.
type Properties struct {
	Properties []models.Property `json:"properties"`
}

func ParsePropertiesJson() string {

	// abrir archivo "configuracion.json"
	manejadorDeArchivo, err := ioutil.ReadFile("db/data.json")
	if err != nil {
		log.Fatal(err)
	}
	// preparar contenedor de las credenciales
	c := Properties{}
	// decodificar el contenido del json sobre la estructura
	err = json.Unmarshal(manejadorDeArchivo, &c)
	if err != nil {
		log.Fatal(err)
	}
	crear_json, _ := json.Marshal(c)

	// Convertimos los datos(bytes) en una cadena e imprimimos el contenido.
	JsonString := string(crear_json)
	return JsonString

}
