package main

import (
	"fmt"
	"project/calculator/fileManager"
	"project/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludePriceJob(fm, taxRate)
		priceJob.Process()
	}
}
