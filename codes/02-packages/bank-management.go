package main

import (
	"bank-app/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.ReadFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("----------------------------------")
	}

	fmt.Println("Welcome to the Great British Bank!")
	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Printf("You have deposited £%.2f. Your new balance is: £%.2f\n", depositAmount, accountBalance)

		} else if choice == 3 {
			fmt.Print("Enter the amount to withdraw: £")
			var withdrawalAmount float64
			fmt.Scanln(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds. Transaction cancelled.")
			} else {
				accountBalance -= withdrawalAmount
				fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
				fmt.Printf("You have withdrawn £%.2f. Your new balance is: £%.2f\n", withdrawalAmount, accountBalance)
			}

		} else if choice == 4 {
			fmt.Println("Thank you for banking with us. Goodbye!")
			break

		} else {
			fmt.Println("Invalid choice. Please try again.")
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)
		}
	}
}
