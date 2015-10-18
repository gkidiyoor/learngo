package datastore

import (
	"log"

	"github.com/abhigupta912/learngo/gorillamux/product"
	"github.com/satori/go.uuid"
)

var Products = make(map[string]product.Product)

func AddNewProduct(prod product.Product) string {
	var id = uuid.NewV4().String()
	log.Printf("Adding Product with Id [%s]: %v", id, prod)
	Products[id] = prod
	return id
}

func ReplaceProduct(id string, newProd product.Product) (product.Product, bool) {
	oldProd, ok := Products[id]
	if ok {
		log.Printf("Replacing Product [%v] with [%v]", oldProd, newProd)
		Products[id] = newProd
	} else {
		log.Printf("No Product with Id [%s] found. Unable to replace.", id)
	}
	return oldProd, ok
}

func RemoveProduct(id string) (product.Product, bool) {
	prod, ok := Products[id]
	if ok {
		log.Printf("Removing Product [%v] with Id [%s]", prod, id)
		delete(Products, id)
	} else {
		log.Printf("No Product with Id [%s] found. Unable to remove.", id)
	}
	return prod, ok
}

func GetProduct(id string) (product.Product, bool) {
	prod, ok := Products[id]
	if ok {
		log.Printf("Retrieving Product with Id [%s]: %v", id, prod)
	} else {
		log.Printf("No Product with Id [%s] found", id)
	}
	return prod, ok
}
