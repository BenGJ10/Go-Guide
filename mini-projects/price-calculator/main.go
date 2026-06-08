package main

import (
	"fmt"
	"price-calculator/filemanager"
	prices "price-calculator/prices"
)

func main() {

	taxRates := []float64{0.01, 0.07, 0.15}

	for _, taxRate := range taxRates {
		fileManager := filemanager.New("input_prices.txt", fmt.Sprintf("output_prices_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPrices(fileManager, taxRate)
		priceJob.ProcessPrices()
	}
}
