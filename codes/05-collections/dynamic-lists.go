package main

import "fmt"

func dynamicLists() {
	prices := []float64{}

	fmt.Println("Initial size of prices: ", len(prices))

	for i := 1; i <= 5; i++ {
		prices = append(prices, float64(i))
	}
	fmt.Println("Prices:", prices)

	fmt.Println("Current size of prices: ", len(prices))

	// Removing the last price
	prices = prices[:len(prices)-1]
	fmt.Println("Prices after removing the last item:", prices)
}
