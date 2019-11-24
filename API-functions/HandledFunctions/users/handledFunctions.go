package users

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

func (h HandledFunctions) GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User
		var user models.User
		rows, _ := db.Query("select * from users")
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&user.ID, &user.UserName, &user.Password)

			users = append(users, user)
		}
		json.NewEncoder(w).Encode(users)
	}
}

func (h HandledFunctions) GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idS := params["id"]
		id, _ := strconv.Atoi(idS)
		var user models.User
		rows := db.QueryRow("select * from users where id=$1", id)
		err := rows.Scan(&user.ID, &user.UserName, &user.Password)
		models.LogFatal(err)
		json.NewEncoder(w).Encode(user)
	}
}
