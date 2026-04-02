package prices

import (
	"fmt"
	"project/calculator/conversion"
	"project/calculator/fileManager"
)

type TaxIncludePriceJob struct {
	IOManager       fileManager.FileManager `json:"-"`
	TaxRate         float64                 `json:"tax_rate"`
	InputPrices     []float64               `json:"input_prices"`
	TaxIncludePrice map[string]string       `json:"tax_include_price"`
}

func (job *TaxIncludePriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
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
	job.TaxIncludePrice = result
	job.IOManager.WriteResult(job)
}

func NewTaxIncludePriceJob(fm fileManager.FileManager, TaxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		IOManager:   fm,
		TaxRate:     TaxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
