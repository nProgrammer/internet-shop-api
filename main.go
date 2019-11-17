package main

import (
	"encoding/json"
	/* LOCAL IMPORTS */
	"internet-shop/models"

	/* GITHUB IMPORTS */
	"github.com/gorilla/mux"

	/* COMPILATOR BUILT-IN IMPORTS */
	"log"
	"net/http"
	"strconv"
)

var products []models.Product

func main() {
	products = append(products,
		models.Product{ID: 1, ProductName: "Python: the terminator", Seller: "Py Python", Price: 1992},
		models.Product{ID: 2, ProductName: "Golang: the terminate of Python", Seller: "Go Gopher", Price: 2010})

	rout := mux.NewRouter()

	rout.HandleFunc("/products", getProducts).Methods("GET")
	rout.HandleFunc("/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", rout))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	log.Println(products)
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	urlMap := mux.Vars(r)
	idS := urlMap["id"]
	id, _ := strconv.Atoi(idS)
	product = products[id]
	log.Println(product)
	json.NewEncoder(w).Encode(product)
}
