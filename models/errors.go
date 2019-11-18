package models

import (
	/* COMPILATOR BUILT-IN IMPORTS */
	"log"
)

func LogFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}
