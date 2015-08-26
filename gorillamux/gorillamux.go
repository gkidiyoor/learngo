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
	Name string `json:"name"`
	Desc string `json:"desc"`
	Qty  int16  `json:"qty"`
}

func (p Product) String() string {
	buffer, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(buffer)
}

func productIdGen() func() string {
	counter := 0
	return func() string {
		counter += 1
		return fmt.Sprintf("P%03d", counter)
	}
}

var (
	portStr  = flag.String("port", "9000", "Server Port")
	products = make(map[string]Product)
	idGen    = productIdGen()
)

func AddNewProduct(prod Product) string {
	var id = idGen()
	log.Printf("Adding Product with Id [%s]: %v", id, prod)
	products[id] = prod
	return id
}

func ReplaceProduct(id string, newProd Product) (Product, bool) {
	oldProd, ok := products[id]
	if ok {
		log.Printf("Replacing Product [%v] with [%v]", oldProd, newProd)
		products[id] = newProd
	} else {
		log.Printf("No Product with Id [%s] found. Unable to replace.", id)
	}
	return oldProd, ok
}

func RemoveProduct(id string) (Product, bool) {
	prod, ok := products[id]
	if ok {
		log.Printf("Removing Product [%v] with Id [%s]", prod, id)
		delete(products, id)
	} else {
		log.Printf("No Product with Id [%s] found. Unable to remove.", id)
	}
	return prod, ok
}

func GetProduct(id string) (Product, bool) {
	prod, ok := products[id]
	if ok {
		log.Printf("Retrieving Product with Id [%s]: %v", id, prod)
	} else {
		log.Printf("No Product with Id [%s] found", id)
	}
	return prod, ok
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Products server")
	fmt.Fprintln(w, "GET /products to get a listing of all Products")
	fmt.Fprintln(w, "GET /product/id to get a listing of Product specified by id")
	fmt.Fprintln(w, "POST /product with Product name, desc and qty to add a new Product")
	fmt.Fprintln(w, "PUT /product/id with Product name, desc and qty to replace existing Product specified by id")
	fmt.Fprintln(w, "DELETE /product/id to remove a Product specified by id")
}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Getting all Products")
	for id, prod := range products {
		fmt.Fprintf(w, "Product Id: %s Product: %v\n", id, prod)
	}
}

func GetProductWithIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	prod, ok := GetProduct(id)
	if ok {
		fmt.Fprintf(w, "Product Id: %s Product: %v\n", id, prod)
	} else {
		http.Error(w, "No such product found\n", http.StatusNotFound)
	}
}

func AddNewProductHandler(w http.ResponseWriter, r *http.Request) {
	prod := Product{}
	err := json.NewDecoder(r.Body).Decode(&prod)

	if err != nil {
		http.Error(w, "Unable to parse request body for Product\n", http.StatusBadRequest)
		return
	}

	id := AddNewProduct(prod)
	fmt.Fprintf(w, "Product [%v] added with Id: %s\n", prod, id)
}

func ReplaceProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	prod := Product{}
	err := json.NewDecoder(r.Body).Decode(&prod)

	if err != nil {
		http.Error(w, "Unable to parse request body for Product\n", http.StatusBadRequest)
		return
	}

	oldProd, ok := ReplaceProduct(id, prod)
	if ok {
		fmt.Fprintf(w, "Product [%v] replaced with new Product [%v]\n", oldProd, prod)
	} else {
		http.Error(w, "No such product found\n", http.StatusNotFound)
	}
}

func RemoveProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	prod, ok := RemoveProduct(id)
	if ok {
		fmt.Fprintf(w, "Product [%v] removed\n", prod)
	} else {
		http.Error(w, "No such product found\n", http.StatusNotFound)
	}
}

func main() {
	Router := mux.NewRouter()
	Router.HandleFunc("/", HomeHandler)

	Router.Path("/products").Methods("GET").HandlerFunc(GetAllProductsHandler)
	Router.PathPrefix("/product/{id}").Methods("GET").HandlerFunc(GetProductWithIdHandler)
	Router.PathPrefix("/product").Headers("Content-Type", "application/json").Methods("POST").HandlerFunc(AddNewProductHandler)
	Router.PathPrefix("/product/{id}").Headers("Content-Type", "application/json").Methods("PUT").HandlerFunc(ReplaceProductHandler)
	Router.PathPrefix("/product/{id}").Methods("DELETE").HandlerFunc(RemoveProductHandler)

	flag.Parse()

	_, portErr := strconv.ParseInt(*portStr, 10, 16)
	if portErr != nil {
		log.Fatalf("Unable to start server on invalid port: %s", *portStr)
	}

	log.Println("Starting server ...")
	log.Printf("Server Listening on Port: %s\n", *portStr)

	log.Fatal(http.ListenAndServe(":"+*portStr, Router))
}
