package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AddUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
}

type User struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	UserName          string    `json:"userName"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"createdAt"`
}

// password not actually implemented yet
type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func NewUser(firstName string, lastName string, userName string, password string) (*User, error) {
	//TODO: format data properly, look into SQL & POSTGRES data types
	//TODO: Find a good system for ID's, how can we expose the ID to the client, should we?

	currentTime := time.Now().UTC()

	encpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         firstName,
		LastName:          lastName,
		CreatedAt:         currentTime,
		UserName:          userName,
		EncryptedPassword: string(encpassword),
	}, nil
}

/*
	Java Springboard
	Brush up on types of testing
	90% testing, cucumber, J unit, selenium, jenkins, brush up on agile
*/
