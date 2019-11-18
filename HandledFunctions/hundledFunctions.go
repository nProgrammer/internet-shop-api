package HandledFunctions

import (
	/* LOCAL IMPORTS */
	"internet-shop/models"

	/* COMPILATOR BUILT-IN IMPORTS */
	"database/sql"
	"encoding/json"
	"net/http"
)

type HandledFunctions struct{}

func (h HandledFunctions) GetProducts(db *sql.DB, products []models.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
