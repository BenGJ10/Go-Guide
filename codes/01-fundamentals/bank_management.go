package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	// If the file doesn't exist or there's an error reading it, we assume a starting balance of £1000.00
	if err != nil {
		return 1000.0, errors.New("Could not read balance file, default balance: £1000.00")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000.0, errors.New("Could not parse balance from file, default balance: £1000.00")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func bankManagement() {
	accountBalance, err := readBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("----------------------------------")
	}

	fmt.Println("Welcome to the Great British Bank!")
	for {
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
			writeBalanceToFile(accountBalance)
			fmt.Printf("You have deposited £%.2f. Your new balance is: £%.2f\n", depositAmount, accountBalance)

		} else if choice == 3 {
			fmt.Print("Enter the amount to withdraw: £")
			var withdrawalAmount float64
			fmt.Scanln(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds. Transaction cancelled.")
			} else {
				accountBalance -= withdrawalAmount
				writeBalanceToFile(accountBalance)
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
