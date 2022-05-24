package main

import (
	"teelgo/blockchain-go/cli"
	"teelgo/blockchain-go/db"
)

func main() {
	db.InitDB()
	defer db.Close()
	cli.Start()
}
