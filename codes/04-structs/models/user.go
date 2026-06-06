package models

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	User
	email string
}

func NewAdmin(firstName, lastName, birthDate, email string) *Admin {
	return &Admin{
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthDate: birthDate,
			createdAt: time.Now(),
		},
		email: email,
	}
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("All fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) OutputUserDetails() {
	fmt.Printf("User Details:\n")
	fmt.Printf("First Name: %s\n", u.firstName)
	fmt.Printf("Last Name: %s\n", u.lastName)
	fmt.Printf("Birth Date: %s\n", u.birthDate)
	fmt.Printf("Created At: %s\n", u.createdAt.Format(time.RFC1123))
}

func (u *User) ClearUserData() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
	u.createdAt = time.Time{}
}
