package prices

import (
	"fmt"
	conversion "price-calculator/conversions"
	iomanager "price-calculator/iomanager"
)

const inputFileName = "input_prices.txt"

type TaxIncludedPrices struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPrices(io iomanager.IOManager, taxRate float64) *TaxIncludedPrices {
	return &TaxIncludedPrices{
		IOManager:   io,
		TaxRate:     taxRate,
		InputPrices: []float64{10.0, 20.0, 30.0},
	}
}

func (job *TaxIncludedPrices) ProcessPrices() {
	job.LoadPrices()

	resultMap := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		resultMap[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = resultMap
	job.IOManager.WriteResult(job)

}

func (job *TaxIncludedPrices) LoadPrices() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}
