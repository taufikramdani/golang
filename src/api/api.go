package main

import (
	"api/bootstrap"
	"encoding/json"
	"log"
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
	log.Println("config" + string(os.PathSeparator) + "config.json")
	bootstrap.LoadJSON("config"+string(os.PathSeparator)+"config.json", config)
	//Try to connect postgresql
	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	// 	DbUser, DbPassword, DbName)
	// db, err := sql.Open("postgres", dbinfo)
	// checkErr(err)
	// defer db.Close()
	// fmt.Println("# Inserting values")

	// var lastInsertID int
	// err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "XXXXXCFCFCF", "2012-12-09").Scan(&lastInsertID)
	// checkErr(err)
	// fmt.Println("last inserted id =", lastInsertID)
	//xx, _ := json.Marshal(config)
	//log.Println("here 2 ", string(xx))
	bootstrap.Connect(config.Database)
}

var config = &configuration{}

type configuration struct {
	Database bootstrap.Info `json:"Database"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
