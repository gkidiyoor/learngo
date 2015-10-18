package middleware

import "github.com/gorilla/mux"

func registerApi() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	router.Path("/persons").Methods("GET").HandlerFunc(getAllPersonsHandler)
	router.PathPrefix("/person/{id}").Methods("GET").HandlerFunc(getPersonWithIdHandler)
	router.Path("/person").Headers("Content-Type", "application/json").Methods("POST").HandlerFunc(addNewPersonHandler)
	router.PathPrefix("/person/{id}").Headers("Content-Type", "application/json").Methods("PUT").HandlerFunc(replacePersonHandler)
	router.PathPrefix("/person/{id}").Methods("DELETE").HandlerFunc(removePersonHandler)

	return router
}
