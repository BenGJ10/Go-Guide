package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	firstName string `filter:"first_name"`
	lastName  string
	birthDate string
}

type Admin struct {
	User
	email string
}

// A function used to create a new user
func newUser(firstName, lastName, birthDate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
	}
}

// A method used to output user details
func (user *User) OutputUserDetails() {
	fmt.Printf("\nFirst Name: %s\n", user.firstName)
	fmt.Printf("Last Name: %s\n", user.lastName)
	fmt.Printf("Birth Date: %s\n", user.birthDate)
}

func (admin *Admin) OutputAdminDetails() {
	admin.OutputUserDetails()
	fmt.Printf("Email: %s\n", admin.email)
}

type randomInt int

func (ri randomInt) isEven() bool {
	return ri%2 == 0
}

func main() {
	initialUser := User{"John", "Doe", "01/01/1990"}
	initialUser.OutputUserDetails()

	secondUser := User{
		firstName: "Jane",
		lastName:  "Smith",
		birthDate: "02/02/1992",
	}
	secondUser.OutputUserDetails()

	thirdUser := newUser("Alice", "Johnson", "03/03/1993")
	thirdUser.OutputUserDetails()

	adminUser := Admin{User{"Bob", "Marly", "04/04/1985"}, "bob@example.com"}
	adminUser.OutputAdminDetails()

	// Generate a random number between 0 and 99
	num := rand.Intn(100)
	var randomNum randomInt = randomInt(num)

	fmt.Printf("\nRandom Number: %d\n", randomNum)
	fmt.Printf("Is the random number even? %t\n", randomNum.isEven())
}
