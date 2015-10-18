package datastore

import "gopkg.in/mgo.v2"

const (
	DBNAME = "test"
	TABLE  = "users"
)

func GetNewMongoSession() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
