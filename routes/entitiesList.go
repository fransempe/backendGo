package routes

import (
	"fmt"
	"net/http"
	//"github.com/fransempe/backendGo/handlers"
)

func EntitiesList(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Listado de entidades")
}
