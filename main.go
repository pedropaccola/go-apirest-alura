package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting REST server")
	HandleRequest()
	fmt.Println("REST server started")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

