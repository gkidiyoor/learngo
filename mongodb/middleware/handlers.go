package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abhigupta912/learngo/mongodb/model"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Person server")
	fmt.Fprintln(w, "GET /persons to get a listing of all Persons")
	fmt.Fprintln(w, "GET /person/id to get details of the Person specified by id")
	fmt.Fprintln(w, "POST /person with Person name, gender and age to add a new Person")
	fmt.Fprintln(w, "PUT /person/id with Person name, gender and age to replace existing Person specified by id")
	fmt.Fprintln(w, "DELETE /person/id to remove a Person specified by id")
}

func getAllPersonsHandler(w http.ResponseWriter, r *http.Request) {
	datastore := GetDataStore(r)
	persons, err := datastore.FindAllPersons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(persons) > 0 {
		fmt.Fprintln(w, "Getting all Persons")
		for _, person := range persons {
			fmt.Fprintf(w, "Person: %v\n", person)
		}
	} else {
		http.Error(w, "No Persons found", http.StatusNotFound)
	}
}

func getPersonWithIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	datastore := GetDataStore(r)
	person, err := datastore.FindPerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Found Person: %v\n", person)
}

func addNewPersonHandler(w http.ResponseWriter, r *http.Request) {
	person := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		http.Error(w, "Unable to parse request body for Person", http.StatusBadRequest)
		return
	}

	datastore := GetDataStore(r)
	id, err := datastore.InsertPerson(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Person %v added with Id: %s\n", person, id)
}

func replacePersonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	person := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		http.Error(w, "Unable to parse request body for Person", http.StatusBadRequest)
		return
	}

	datastore := GetDataStore(r)
	err = datastore.UpdatePerson(id, person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func removePersonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	datastore := GetDataStore(r)
	err := datastore.RemovePerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
