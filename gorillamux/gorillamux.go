package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/abhigupta912/learngo/gorillamux/datastore"
	"github.com/abhigupta912/learngo/gorillamux/product"
	"github.com/gorilla/mux"
)

var portStr = flag.String("port", "9000", "Server Port")

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Products server")
	fmt.Fprintln(w, "GET /products to get a listing of all Products")
	fmt.Fprintln(w, "GET /product/id to get a listing of Product specified by id")
	fmt.Fprintln(w, "POST /product with Product name, desc and qty to add a new Product")
	fmt.Fprintln(w, "PUT /product/id with Product name, desc and qty to replace existing Product specified by id")
	fmt.Fprintln(w, "DELETE /product/id to remove a Product specified by id")
}

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	if len(datastore.Products) > 0 {
		fmt.Fprintln(w, "Getting all Products")
		for id, prod := range datastore.Products {
			fmt.Fprintf(w, "Product Id: %s Product: %v\n", id, prod)
		}
	} else {
		fmt.Fprintln(w, "No Products found")
	}
}

func getProductWithIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	prod, ok := datastore.GetProduct(id)
	if ok {
		fmt.Fprintf(w, "Product Id: %s Product: %v\n", id, prod)
	} else {
		http.Error(w, "No such product found", http.StatusNotFound)
	}
}

func addNewProductHandler(w http.ResponseWriter, r *http.Request) {
	prod := product.Product{}
	err := json.NewDecoder(r.Body).Decode(&prod)

	if err != nil {
		http.Error(w, "Unable to parse request body for Product", http.StatusBadRequest)
		return
	}

	id := datastore.AddNewProduct(prod)
	fmt.Fprintf(w, "Product [%v] added with Id: %s\n", prod, id)
}

func replaceProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	prod := product.Product{}
	err := json.NewDecoder(r.Body).Decode(&prod)

	if err != nil {
		http.Error(w, "Unable to parse request body for Product", http.StatusBadRequest)
		return
	}

	oldProd, ok := datastore.ReplaceProduct(id, prod)
	if ok {
		fmt.Fprintf(w, "Product [%v] replaced with new Product [%v]\n", oldProd, prod)
	} else {
		http.Error(w, "No such product found", http.StatusNotFound)
	}
}

func removeProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	prod, ok := datastore.RemoveProduct(id)
	if ok {
		fmt.Fprintf(w, "Product [%v] removed\n", prod)
	} else {
		http.Error(w, "No such product found", http.StatusNotFound)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	router.Path("/products").Methods("GET").HandlerFunc(getAllProductsHandler)
	router.PathPrefix("/product/{id}").Methods("GET").HandlerFunc(getProductWithIdHandler)
	router.Path("/product").Headers("Content-Type", "application/json").Methods("POST").HandlerFunc(addNewProductHandler)
	router.PathPrefix("/product/{id}").Headers("Content-Type", "application/json").Methods("PUT").HandlerFunc(replaceProductHandler)
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
