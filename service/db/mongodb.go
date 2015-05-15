package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoPerson struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Username  string        `bson:"username,omitempty"`
	FirstName string        `bson:"first_name,omitempty"`
	LastName  string        `bson:"last_name,omitempty"`
	Age       int32         `bson:"age,omitempty"`
	Hobbies   []string      `bson:"hobbies,omitempty"`
	Team      *mongoTeam    `bson:"team,omitempty"`
}

type mongoTeam struct {
	Name   string `bson:"name,omitempty"`
	Active bool   `bson:"active,omitempty"`
}

type MongoDB struct {
	session *mgo.Session
}

func NewMongoDB(mongoURL string) (DB, error) {
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return &MongoDB{session}, nil
}

func (m *MongoDB) Ping() error {
	return m.session.Ping()
}

func (m *MongoDB) Close() error {
	m.session.Close()
	return nil
}

func (m *MongoDB) GetPerson(username string) (*Person, error) {
	session := m.session.Copy()
	defer session.Close()
	query := session.DB("").C("people").Find(bson.M{"username": username})
	var result mongoPerson
	if err := query.One(&result); err != nil {
		return nil, fmt.Errorf("Error finding person: %s", err)
	}
	person := &Person{
		Id:        result.Id.Hex(),
		Username:  result.Username,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Age:       result.Age,
		Hobbies:   result.Hobbies,
	}
	if result.Team != nil {
		person.Team = &Team{
			Name:   result.Team.Name,
			Active: result.Team.Active,
		}
	}
	return person, nil
}
