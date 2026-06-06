package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.00

	fmt.Println("Welcome to the Great British Bank!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check your balance")
	fmt.Println("2. Make a deposit")
	fmt.Println("3. Make a withdrawal")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Printf("Your current balance is: £%.2f\n", accountBalance)
	} else if choice == 2 {
		fmt.Print("Enter the amount to deposit: £")
		var depositAmount float64
		fmt.Scanln(&depositAmount)
		accountBalance += depositAmount
		fmt.Printf("You have deposited £%.2f. Your new balance is: £%.2f\n", depositAmount, accountBalance)
	} else if choice == 3 {
		fmt.Print("Enter the amount to withdraw: £")
		var withdrawalAmount float64
		fmt.Scanln(&withdrawalAmount)
		if withdrawalAmount > accountBalance {
			fmt.Println("Insufficient funds. Transaction cancelled.")
		} else {
			accountBalance -= withdrawalAmount
			fmt.Printf("You have withdrawn £%.2f. Your new balance is: £%.2f\n", withdrawalAmount, accountBalance)
		}
	} else if choice == 4 {
		fmt.Println("Thank you for banking with us. Goodbye!")
		return
	} else {
		fmt.Println("Invalid choice. Please try again.")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
	}
}
