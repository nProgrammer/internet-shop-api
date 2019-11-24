package main

import (
	"github.com/subosito/gotenv"
	"internet-shop/API-functions/cfg"
)

func init() {
	gotenv.Load()
}

func main() {
	cfg.StartApi()
}
