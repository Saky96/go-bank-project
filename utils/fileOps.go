package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetTextFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteTextToFile(text float64, fileName string) {
	textValue := fmt.Sprint(text)
	os.WriteFile(fileName, []byte(textValue), 0644)
}
