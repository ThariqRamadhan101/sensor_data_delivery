package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getAll", getAll).Methods("GET")
	router.HandleFunc("/registration", getRegistration).Methods("GET")
	router.HandleFunc("/question/{id}", getQuestion).Methods("GET")
	router.HandleFunc("/data", getData).Methods("POST")
	router.HandleFunc("/reset/{id}", getReset).Methods("GET")
	router.HandleFunc("/sync/{id}", getSync).Methods("GET")
	router.HandleFunc("/history/{id}/{n}", getHistory).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}