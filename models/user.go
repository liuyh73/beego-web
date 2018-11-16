package models

type User struct {
	Username  string
	Password  string
	Email     string
	Telephone string
	Id        string
}

var Users = make([]User, 0)