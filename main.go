package main

import (
	"github.com/subosito/gotenv"
	"internet-shop/API-functions"
)

func init() {
	gotenv.Load()
}

func main() {
	API_functions.StartApi()
}
