package products

import (
	/* GITHUB IMPORTS */
	"github.com/gorilla/mux"

	/* LOCAL IMPORTS */
	"internet-shop/models"

	/* COMPILATOR BUILT-IN IMPORTS */
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type HandledFunctions struct{}

func (h HandledFunctions) GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var products []models.Product
		var product models.Product
		rows, _ := db.Query("select * from products")
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&product.ID, &product.ProductName, &product.Seller, &product.Price)

			products = append(products, product)
		}
		json.NewEncoder(w).Encode(products)
	}
}

func (h HandledFunctions) GetProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idS := params["id"]
		id, _ := strconv.Atoi(idS)
		var product models.Product
		rows := db.QueryRow("select * from products where id=$1", id)
		err := rows.Scan(&product.ID, &product.ProductName, &product.Seller, &product.Price)
		models.LogFatal(err)
		json.NewEncoder(w).Encode(product)
	}
}

func (h HandledFunctions) DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idS := params["id"]
		id, _ := strconv.Atoi(idS)
		result, err := db.Exec("delete from products where id=$1", id)
		models.LogFatal(err)
		rowsUpdated, err := result.RowsAffected()
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}
