package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedropaccola/go-restapi-alura/database"
	"github.com/pedropaccola/go-restapi-alura/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func People(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    p := []models.Person{} 
    database.DB.Find(&p)
	json.NewEncoder(w).Encode(models.People)
}

func Person(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
    p := models.Person{}
    database.DB.First(&p, id)

    json.NewEncoder(w).Encode(p)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
    p := models.Person{}
    json.NewDecoder(r.Body).Decode(&p)
    database.DB.Create(&p)
    json.NewEncoder(w).Encode(p)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
    p := models.Person{}
    database.DB.Delete(&p, id)
    json.NewEncoder(w).Encode(p)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
    p := models.Person{}
    database.DB.First(&p, id)
    json.NewDecoder(r.Body).Decode(&p)
    database.DB.Save(&p)
    json.NewEncoder(w).Encode(p)
}
