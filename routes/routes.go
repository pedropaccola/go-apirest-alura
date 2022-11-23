package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedropaccola/go-restapi-alura/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/people/{id}", controllers.Person)
	r.HandleFunc("/api/people", controllers.People)

	log.Fatal(http.ListenAndServe(":8000", r))
}
