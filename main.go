package main

import (
	"database/sql"
	"kitap-api/base"
	"kitap-api/mod"
	"log"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []mod.Book
var db *sql.DB

func init() {
	gotenv.Load()
}
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db = base.ConnectDB()
	router := mux.NewRouter()
	handlerss := hundl.Handlers{}

	router.HandleFunc("/books", hundlers.GetBooks(db)).Methods("GET")
	/*
		router.HandleFunc("/books/{id}", getBook).Methods("GET")
		router.HandleFunc("/books", addBook).Methods("POST")
		router.HandleFunc("/books", updateBook).Methods("PUT")
		router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

		log.Fatal(http.ListenAndServe(":9000", router))
	*/
}
