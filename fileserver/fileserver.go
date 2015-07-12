package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	portStr = flag.String("port", "9000", "Server Port")
	dirName = flag.String("dir", ".", "Root Directory")
)

func main() {
	flag.Parse()

	_, portErr := strconv.ParseInt(*portStr, 10, 16)
	if portErr != nil {
		log.Fatalf("Unable to start server on invalid port: %s", *portStr)
	}

	fileInfo, fileErr := os.Stat(*dirName)
	if fileErr != nil {
		log.Fatalf("Unable to serve files from : %s", *dirName)
		log.Fatal(fileErr)
	}

	if !fileInfo.IsDir() {
		log.Fatalf("Error: %s is not a directory", *dirName)
	}

	log.Println("Starting server ...")
	log.Printf("Server Listening on Port: %s\n", *portStr)
	log.Printf("Serving Directory: %s\n", *dirName)

	httpErr := http.ListenAndServe(":"+*portStr, http.FileServer(http.Dir(*dirName)))
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
