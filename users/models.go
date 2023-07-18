package users

import (
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(firstName string, lastName string, userName string) *User {
	//TODO: format data properly, look into SQL & POSTGRES data types
	//TODO: Find a good system for ID's, how can we expose the ID to the client, should we?

	currentTime := time.Now().UTC()
	//ID: rand.Intn(10000),
	return &User{FirstName: firstName, LastName: lastName, CreatedAt: currentTime, UserName: "working on it"}
}
