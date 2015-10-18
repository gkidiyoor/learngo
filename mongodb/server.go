package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/abhigupta912/learngo/mongodb/middleware"
)

var portStr = flag.String("port", "9000", "Server Port")

func main() {
	flag.Parse()

	_, portErr := strconv.ParseInt(*portStr, 10, 16)
	if portErr != nil {
		log.Fatalf("Unable to start server on invalid port: %s", *portStr)
	}

	log.Println("Starting server ...")
	log.Printf("Server Listening on Port: %s\n", *portStr)

	middleware.NegroniChain(*portStr)
}
