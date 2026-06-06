package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000.0, errors.New("Could not read file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000.0, errors.New("Could not parse stored value from file")
	}
	return value, nil
}

func WriteFloatToFile(filename string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}
