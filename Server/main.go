package main

import (
	"log"
	"net/http"
<<<<<<< HEAD
	_ "time"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getAll", getAll).Methods("GET")
	router.HandleFunc("/registration", getRegistration).Methods("POST")
=======

	"github.com/gorilla/mux"
)


func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getAll", getAll).Methods("GET")
	router.HandleFunc("/registration", getRegistration).Methods("GET")
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
	router.HandleFunc("/question/{id}", getQuestion).Methods("GET")
	router.HandleFunc("/data", getData).Methods("POST")
	router.HandleFunc("/reset/{id}", getReset).Methods("GET")
	router.HandleFunc("/sync/{id}", getSync).Methods("GET")
	router.HandleFunc("/history/{id}/{n}", getHistory).Methods("GET")
<<<<<<< HEAD
	log.Fatal(http.ListenAndServe(":8080", limit(router)))
}
=======
	log.Fatal(http.ListenAndServe(":8080", router))
}
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
