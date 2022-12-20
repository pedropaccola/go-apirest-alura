package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/pedropaccola/go-restapi-alura/controllers"
	"github.com/pedropaccola/go-restapi-alura/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
    r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/people/{id}", handlePeopleId)
	r.HandleFunc("/api/people", handlePeople)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

func handlePeopleId(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "Get":
        controllers.Person(w, r)
    case "Delete":
        controllers.DeletePerson(w, r)
    case "Put":
        controllers.EditPerson(w, r)
    }
}

func handlePeople(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "Get":
        controllers.People(w, r)
    case "Post":
        controllers.CreatePerson(w, r)
    }
}
