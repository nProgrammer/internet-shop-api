package cfg

import (
	/* LOCAL IMPORTS */
	"internet-shop/HandledFunctions/products"
	"internet-shop/HandledFunctions/users"

	/* GITHUB IMPORTS */
	"github.com/gorilla/mux"

	/* COMPILATOR BUILT-IN IMPORTS */
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func StartApi() {
	handlerP := products.HandledFunctions{}
	handlerU := users.HandledFunctions{}

	db = ConnectDB(db)

	log.Println("server is running on port :8000")

	rout := mux.NewRouter()

	/* PRODUCTS HANDLERS */
	rout.HandleFunc("/products", handlerP.GetProducts(db)).Methods("GET")
	rout.HandleFunc("/products/{id}", handlerP.GetProduct(db)).Methods("GET")
	rout.HandleFunc("/products/{id}", handlerP.DeleteProduct(db)).Methods("DELETE")

	/* USERS HANDLERS */
	rout.HandleFunc("/users", handlerU.GetUsers(db)).Methods("GET")
	rout.HandleFunc("/users/{id}", handlerU.GetUser(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", rout))
}
