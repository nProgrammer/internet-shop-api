package cfg

import (
	/* LOCAL IMPORTS */
	"internet-shop/models"

	/* GITHUB IMPORTS */
	"github.com/lib/pq"

	/* COMPILATOR BUILT-IN IMPORTS */
	"database/sql"
	"log"
	"os"
)

func ConnectDB(db *sql.DB) *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	models.LogFatal(err)
	log.Println(pgUrl)

	db, err = sql.Open("postgres", pgUrl)
	models.LogFatal(err)

	err = db.Ping()
	models.LogFatal(err)

	return db
}
