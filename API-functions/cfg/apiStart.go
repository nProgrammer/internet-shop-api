package cfg

import (
	/* LOCAL IMPORTS */
	handle "internet-shop/API-functions"
	products2 "internet-shop/API-functions/HandledFunctions/products"
	users2 "internet-shop/API-functions/HandledFunctions/users"

	/* COMPILATOR BUILT-IN IMPORTS */
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func StartApi() {
	handlerP := products2.HandledFunctions{}
	handlerU := users2.HandledFunctions{}

	db = ConnectDB(db)

	rout := handle.RouterFunc(handlerP, handlerU, db)

	log.Fatal(http.ListenAndServe(":8000", rout))

	log.Println("server is running on port :8000")
}
