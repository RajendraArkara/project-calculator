package prices

import (
	"fmt"
	"project/calculator/conversion"
	"project/calculator/fileManager"
)

type TaxIncludePriceJob struct {
	TaxRate         float64
	InputPrices     []float64
	TaxIncludePrice map[string]float64
}

func (job *TaxIncludePriceJob) LoadData() {
	lines, err := fileManager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludePriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludePrice)
	}
	fmt.Println(result)
}

func NewTaxIncludePriceJob(TaxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		TaxRate:     TaxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
