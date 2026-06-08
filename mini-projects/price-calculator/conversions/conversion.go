package conversions

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := []float64{}
	for _, line := range strings {
		floatVal, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}
