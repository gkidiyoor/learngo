package model

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
}

type PersonDocument struct {
	Person
	Id bson.ObjectId `json:"id" bson:"_id"`
}

func (p Person) String() string {
	buffer, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(buffer)
}
