package prices

import (
	"fmt"
	"project/calculator/conversion"
	iomanager "project/calculator/ioManager"
)

type TaxIncludePriceJob struct {
	IOManager       iomanager.IOManager `json:"-"`
	TaxRate         float64             `json:"tax_rate"`
	InputPrices     []float64           `json:"input_prices"`
	TaxIncludePrice map[string]string   `json:"tax_include_price"`
}

func (job *TaxIncludePriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludePriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludePrice)
	}
	job.TaxIncludePrice = result
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludePriceJob(iom iomanager.IOManager, TaxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		IOManager:   iom,
		TaxRate:     TaxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
