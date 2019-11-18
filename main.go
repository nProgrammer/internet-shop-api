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
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var products []models.Product
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	hundler := HandledFunctions.HandledFunctions{}

	db = cfg.ConnectDB(db)

	log.Println("server is running on port :8000")

	rout := mux.NewRouter()
	rout.HandleFunc("/products", hundler.GetProducts(db, products)).Methods("GET")
	//rout.HandleFunc("/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", rout))
}
