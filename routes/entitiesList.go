package routes

import (
	"fmt"
	"net/http"
)

func EntitiesList(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Estas son las entidades")
}
