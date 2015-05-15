package db

type Person struct {
	Id        string
	Username  string
	FirstName string
	LastName  string
	Age       int32
	Hobbies   []string
	Team      *Team
}

type Team struct {
	Name   string
	Active bool
}

type DB interface {
	Ping() error
	Close() error
	GetPerson(string) (*Person, error)
}
