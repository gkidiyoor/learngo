package middleware

import (
	"net/http"

	"github.com/abhigupta912/learngo/mongodb/datastore"
	"github.com/gorilla/context"
)

const dataStoreKey int = 0

func GetDataStore(req *http.Request) *datastore.DataStore {
	if ds := context.Get(req, dataStoreKey); ds != nil {
		return ds.(*datastore.DataStore)
	}

	return nil
}

func SetDataStore(req *http.Request, datastore *datastore.DataStore) {
	context.Set(req, dataStoreKey, datastore)
}

func ClearDataStore(req *http.Request) {
	context.Clear(req)
}
