package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pedropaccola/go-restapi-alura/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func People(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.People)
}

func Person(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, person := range models.People {
		if strconv.Itoa(person.ID) == id {
			json.NewEncoder(w).Encode(person)
		}
	}
}
