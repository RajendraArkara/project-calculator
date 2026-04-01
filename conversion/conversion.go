package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, StringVal := range strings {
		floatVal, err := strconv.ParseFloat(StringVal, 64)
		if err != nil {
			return nil, errors.New("Error")
		}

		floats = append(floats, floatVal)
	}
	return floats, nil
}
