package middleware

import (
	"net/http"

	"github.com/abhigupta912/learngo/mongodb/datastore"
	"github.com/codegangsta/negroni"
)

func mongoMiddleware() negroni.HandlerFunc {
	session, err := datastore.GetNewMongoSession()

	if err != nil {
		panic(err)
	}

	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		sessionClone := session.Clone()
		defer sessionClone.Close()

		dstore := datastore.NewDataStore(sessionClone)
		SetDataStore(req, dstore)
		next(rw, req)
		ClearDataStore(req)
	})
}

func NegroniChain(port string) {
	n := negroni.Classic()
	n.Use(mongoMiddleware())
	n.UseHandler(registerApi())
	n.Run(":" + port)
}
