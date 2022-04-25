package main

import (
	"tetgo/tetgocoin/clii"
	"tetgo/tetgocoin/db"
)

func main() {
	defer db.Close()
	clii.Start()
}
