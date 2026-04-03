package main

import (
	"fmt"
	"project/calculator/cmdmanager"
	"project/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := fileManager.New("pricessss.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludePriceJob(cmdm, taxRate)
		priceJob.Process()

		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}
}
