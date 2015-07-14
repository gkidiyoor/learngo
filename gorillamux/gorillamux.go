package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	Id   string "json:id"
	Desc string "json:desc"
	Qty  int16  "json:qty"
}

var (
	portStr  = flag.String("port", "9000", "Server Port")
	products = make(map[string]Product)
)

func addReplaceProduct(id string, desc string, qty int16) {
	products[id] = Product{id, desc, qty}
}

func removeProduct(id string) {
	delete(products, id)
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Welcome to Products server")
	fmt.Fprintln(rw, "GET /products to get a listing of all Products")
	fmt.Fprintln(rw, "GET /product/id to get a listing of Product specified by id")
	fmt.Fprintln(rw, "POST /product/id with desc and qty to add a new Product specified by id")
	fmt.Fprintln(rw, "DELETE /product/id to remove a Product specified by id")
}

func productsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(rw)

	for _, prod := range products {
		encoder.Encode(prod)
	}
}

func getProductHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	rw.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(rw)

	if prod, ok := products[id]; ok {
		encoder.Encode(prod)
		return
	}

	http.Error(rw, "No such product found", http.StatusNotFound)
}

func addReplaceProductHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	prod := Product{Id: id}
	err := json.NewDecoder(r.Body).Decode(&prod)

	if err != nil {
		log.Printf("Error decoding request body", err)
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		return
	}

	addReplaceProduct(id, prod.Desc, prod.Qty)
}

func removeProductHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	removeProduct(id)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	productsApi := router.Path("/products").Subrouter()
	productsApi.Methods("GET").HandlerFunc(productsHandler)

	router.PathPrefix("/product/{id}").Methods("GET").HandlerFunc(getProductHandler)
	router.PathPrefix("/product/{id}").Headers("Content-Type", "application/json").Methods("POST").HandlerFunc(addReplaceProductHandler)
	router.PathPrefix("/product/{id}").Methods("DELETE").HandlerFunc(removeProductHandler)

	flag.Parse()

	_, portErr := strconv.ParseInt(*portStr, 10, 16)
	if portErr != nil {
		log.Fatalf("Unable to start server on invalid port: %s", *portStr)
	}

	log.Println("Starting server ...")
	log.Printf("Server Listening on Port: %s\n", *portStr)

	log.Fatal(http.ListenAndServe(":"+*portStr, router))
}
