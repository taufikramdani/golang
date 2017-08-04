package main

import (
	"api/bootstrap"
	"encoding/json"
	"os"
)

//User struct
type User struct {
	ID   int
	Name string
	//Address Address
}

//UserService interface
type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}

// main func
func main() {

	bootstrap.LoadJSON("config"+string(os.PathSeparator)+"config.json", config)
	bootstrap.Connect(config.Database)
}

var config = &configuration{}

type configuration struct {
	Database bootstrap.Info `json:"Database"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
