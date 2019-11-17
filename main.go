package main

import (
	/* LOCAL IMPORTS */
	"./models"

	/* GITHUB IMPORTS */
	"github.com/gorilla/mux"

	/* COMPILATOR BUILT-IN IMPORTS */
	"log"
	"net/http"
)

var products []models.Product

func main() {
	products = append(products,
		models.Product{ID: 1, Name: "Python: the terminator", FilmDirector: "Py Python", Year: "1992"},
		models.Product{ID: 2, Name: "Golang: the terminate of Python", FilmDirector: "Go Gopher", Year: "2010"})

	rout := mux.NewRouter()

	rout.HandleFunc("/products", getProducts).Methods("GET")
	rout.HandleFunc("/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", rout))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("All products")
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	urlMap := mux.Vars(r)
	id := urlMap["id"]
	log.Printf("Product with id=%s", id)
}
