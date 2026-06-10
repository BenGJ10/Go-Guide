package main

import (
	"fmt"
	"price-calculator/filemanager"
	prices "price-calculator/prices"
)

func main() {

	taxRates := []float64{0.01, 0.07, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChans[idx] = make(chan bool)
		fileManager := filemanager.New("input_prices.txt", fmt.Sprintf("output_prices_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPrices(fileManager, taxRate)
		go priceJob.ProcessPrices(doneChans[idx])
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
