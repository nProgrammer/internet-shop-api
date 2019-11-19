package main

import (
	/* LOCAL IMPORTS */
	"internet-shop/HandledFunctions"
	"internet-shop/cfg"
	"internet-shop/models"

	/* GITHUB IMPORTS */
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"

	/* COMPILATOR BUILT-IN IMPORTS */
	"database/sql"
	"log"
	"net/http"
)

var products []models.Product
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	handler := HandledFunctions.HandledFunctions{}

	db = cfg.ConnectDB(db)

	log.Println("server is running on port :8000")

	rout := mux.NewRouter()
	rout.HandleFunc("/products", handler.GetProducts(db, products)).Methods("GET")
	rout.HandleFunc("/products/{id}", handler.GetProduct(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", rout))
}
