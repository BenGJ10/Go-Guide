package main

import "fmt"

func profitCalculator() {

	revenue := getUserInput("Enter total revenue: ")
	expenses := getUserInput("Enter total expenses: ")
	taxRate := getUserInput("Enter tax rate (in %): ")

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

	// Custom function to print a message
	cout("Profit calculation completed successfully.")
}

func cout(text string) {
	// A utility function to print text, similar to C++'s cout
	fmt.Println("\nBen says:", text)
}

func getUserInput(prompt string) (userInput float64) {
	fmt.Print(prompt)
	fmt.Scanln(&userInput)
	return userInput
}

func calculateValues(revenue, expenses, taxRate float64) (earningsBeforeTax, profit, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - taxRate/100)
	ratio = profit / revenue
	return earningsBeforeTax, profit, ratio
}
