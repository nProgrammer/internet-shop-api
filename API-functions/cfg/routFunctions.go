package cfg

import (
	"database/sql"
	"github.com/gorilla/mux"
	products2 "internet-shop/API-functions/HandledFunctions/products"
	users2 "internet-shop/API-functions/HandledFunctions/users"
)

func RouterFunc(handlerP products2.HandledFunctions, handlerU users2.HandledFunctions, db *sql.DB) *mux.Router {
	rout := mux.NewRouter()

	/* PRODUCTS HANDLERS */
	rout.HandleFunc("/products", handlerP.GetProducts(db)).Methods("GET")
	rout.HandleFunc("/products/{id}", handlerP.GetProduct(db)).Methods("GET")
	rout.HandleFunc("/products/{id}", handlerP.DeleteProduct(db)).Methods("DELETE")

	/* USERS HANDLERS */
	rout.HandleFunc("/users", handlerU.GetUsers(db)).Methods("GET")
	rout.HandleFunc("/users/{id}", handlerU.GetUser(db)).Methods("GET")

	return rout
}
