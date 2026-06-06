package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func investmentCalculator() {
	var initialInvestment, years float64 = 1000, 10
	expectedReturnRate := 5.5

	fmt.Print("Enter the initial investment amount: ")
	fmt.Scanln(&initialInvestment)

	fmt.Print("Enter the number of years for the investment: ")
	fmt.Scanln(&years)

	futureValue, futureRealValue := calculateFutureValue(initialInvestment, expectedReturnRate, years)

	fmt.Printf("Future Value of Investment: %.2f\n", futureValue)
	fmt.Printf("Future Real Value of Investment: %.2f\n", futureRealValue)
}

func calculateFutureValue(initialInvestment, expectedReturnRate, years float64) (futureValue float64, realFutureValue float64) {
	futureValue = initialInvestment * math.Pow(1+expectedReturnRate/100, years)
	realFutureValue = futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, realFutureValue
}
