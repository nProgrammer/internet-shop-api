package API_functions

import (
	products2 "internet-shop/API-functions/HandledFunctions/products"
	users2 "internet-shop/API-functions/HandledFunctions/users"
	"internet-shop/API-functions/cfg"

	/* COMPILATOR BUILT-IN IMPORTS */
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func StartApi() {
	handlerP := products2.HandledFunctions{}
	handlerU := users2.HandledFunctions{}

	db = cfg.ConnectDB(db)

	rout := cfg.RouterFunc(handlerP, handlerU, db)

	log.Fatal(http.ListenAndServe(":8000", rout))

	log.Println("server is running on port :8000")
}
