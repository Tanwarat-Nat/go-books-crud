package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8080"

var books []Book

func main() {
	r := mux.NewRouter()

	books = append(books, Book{
		ID:      "1",
		Isbn:    "438227",
		Title:   "Books One",
		Authors: &Authors{FirstName: "J.K.", LastName: "Rowling"}})

	books = append(books, Book{
		ID:      "2",
		Isbn:    "454555",
		Title:   "Books two",
		Authors: &Authors{FirstName: "Mark", LastName: "Twain"}})

	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books{id}", Updatebooks).Methods("PUT")
	r.HandleFunc("/books{id}", DeleteBooks).Methods("DELETE")

	fmt.Printf("Starting server on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
