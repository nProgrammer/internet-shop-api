package main

import (
	"github.com/subosito/gotenv"
	"internet-shop/cfg"
)

func init() {
	gotenv.Load()
}

func main() {
	cfg.StartApi()
}
