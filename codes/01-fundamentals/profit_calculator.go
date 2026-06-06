package main

import (
	"errors"
	"fmt"
	"os"
)

func profitCalculator() {

	revenue, revErr := getUserInput("Enter total revenue: ")
	expenses, expErr := getUserInput("Enter total expenses: ")
	taxRate, taxErr := getUserInput("Enter tax rate (in %): ")

	if revErr != nil || expErr != nil || taxErr != nil {
		fmt.Println("ERROR", expErr)
		return
	}

	// Calculating using intermediate variables for clarity
	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate/100)
	// ratio := profit / revenue

	// Using a function to calculate all values at once
	earningsBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

	// We can use fmt.Sprintf to format the earnings before tax and then print it separately
	formattedEarnings := fmt.Sprintf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Print(formattedEarnings)

	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Profit to Revenue Ratio: %.2f%%\n", ratio*100)

	storeResults(earningsBeforeTax, profit, ratio)

	// Custom function to print a message
	cout("Profit calculation completed successfully.")
}

func storeResults(earningsBeforeTax, profit, ratio float64) {
	result := fmt.Sprintf("Earnings Before Tax: %.2f\nProfit: %.2f\nProfit to Revenue Ratio: %.2f%%\n", earningsBeforeTax, profit, ratio*100)
	os.WriteFile("result.txt", []byte(result), 0644)
}

func cout(text string) {
	// A utility function to print text, similar to C++'s cout
	fmt.Println("\nBen says:", text)
}

func getUserInput(prompt string) (userInput float64, err error) {
	fmt.Print(prompt)
	fmt.Scanln(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Input must be a positive number")
	}
	return userInput, nil
}

func calculateValues(revenue, expenses, taxRate float64) (earningsBeforeTax, profit, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - taxRate/100)
	ratio = profit / revenue
	return earningsBeforeTax, profit, ratio
}
