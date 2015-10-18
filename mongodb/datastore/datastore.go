package datastore

import (
	"errors"

	"github.com/abhigupta912/learngo/mongodb/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DataStore struct {
	session *mgo.Session
}

func NewDataStore(session *mgo.Session) *DataStore {
	return &DataStore{session}
}

func (datastore *DataStore) InsertPerson(person *model.Person) (string, error) {
	personDocument := model.PersonDocument{}
	personDocument.Id = bson.NewObjectId()
	personDocument.Name = person.Name
	personDocument.Gender = person.Gender
	personDocument.Age = person.Age

	err := datastore.session.DB(DBNAME).C(TABLE).Insert(personDocument)
	if err != nil {
		return "", err
	} else {
		return personDocument.Id.String(), nil
	}
}

func (datastore *DataStore) FindPerson(id string) (model.Person, error) {
	person := model.Person{}
	personDocument := &model.PersonDocument{}

	if !bson.IsObjectIdHex(id) {
		return person, errors.New("Invalid Id")
	}

	personId := bson.ObjectIdHex(id)
	err := datastore.session.DB(DBNAME).C(TABLE).FindId(personId).One(personDocument)
	if err != nil {
		return person, err
	}

	person.Name = personDocument.Name
	person.Gender = personDocument.Gender
	person.Age = personDocument.Age
	return person, nil
}

func (datastore *DataStore) FindAllPersons() ([]model.Person, error) {
	persons := []model.Person{}
	personDocuments := []model.PersonDocument{}

	err := datastore.session.DB(DBNAME).C(TABLE).Find(nil).All(&personDocuments)
	if err != nil {
		return persons, err
	}

	for _, personDocument := range personDocuments {
		person := model.Person{}
		person.Name = personDocument.Name
		person.Gender = personDocument.Gender
		person.Age = personDocument.Age
		persons = append(persons, person)
	}

	return persons, nil
}

func (datastore *DataStore) UpdatePerson(id string, person model.Person) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Id")
	}

	personId := bson.ObjectIdHex(id)
	personDocument := model.PersonDocument{}
	personDocument.Id = personId
	personDocument.Name = person.Name
	personDocument.Gender = person.Gender
	personDocument.Age = person.Age

	return datastore.session.DB(DBNAME).C(TABLE).UpdateId(personId, personDocument)
}

func (datastore *DataStore) RemovePerson(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Id")
	}

	personId := bson.ObjectIdHex(id)
	return datastore.session.DB(DBNAME).C(TABLE).RemoveId(personId)
}
