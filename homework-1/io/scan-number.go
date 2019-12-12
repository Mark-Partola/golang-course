package io

import (
	"errors"
	"fmt"
	"strconv"
)

// ScanNumber - Request CLI input for number
func ScanNumber() (float64, error) {
	var input string
	fmt.Scan(&input)

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("Cannot parse input")
	}

	return num, nil
}
