package clii

import (
	"flag"
	"fmt"
	"os"
	"tetgo/tetgocoin/explorer"
	"tetgo/tetgocoin/restt"
)

func usage() {
	fmt.Printf("Welcome to tetgo coin\n\n")
	fmt.Printf("Please use the following flags\n\n")
	fmt.Printf("-port: Set the PORT of the server\n")
	fmt.Printf("-mode: Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()

	switch *mode {
	case "rest":
		restt.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
