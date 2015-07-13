package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var portStr = flag.String("port", "9000", "Server Port")

func echohandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(rw, "Welcome to simple echo server")
		fmt.Fprintln(rw, "POST a request to receive message back")
		return
	}

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Println("Error reading request body")
		return
	}

	bodyStr := string(body)
	log.Printf("Request body: %s", bodyStr)
	fmt.Fprintln(rw, bodyStr)
}

func main() {
	flag.Parse()

	_, portErr := strconv.ParseInt(*portStr, 10, 16)
	if portErr != nil {
		log.Fatalf("Unable to start server on invalid port: %s", *portStr)
	}

	log.Println("Starting server ...")
	log.Printf("Server Listening on Port: %s\n", *portStr)

	http.HandleFunc("/", echohandler)
	log.Fatal(http.ListenAndServe(":"+*portStr, nil))
}
