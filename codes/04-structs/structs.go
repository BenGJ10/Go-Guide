package main

import (
	"fmt"
	mod "user-app/models"
)

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthDate := getUserData("Enter your birth date (MM/DD/YYYY): ")

	var newUser *mod.User
	newUser, err := mod.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println("❌ Error creating user:", err)
		return
	}

	newUser.OutputUserDetails()
	newUser.ClearUserData()
	fmt.Println("\nClearing user data...")
	newUser.OutputUserDetails()

	fmt.Println("\nCreating admin data...")
	var adminUser *mod.Admin
	adminUser = mod.NewAdmin("Ben", "Gregory", "01/01/1990", "admin@example.com")
	adminUser.OutputUserDetails()
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
