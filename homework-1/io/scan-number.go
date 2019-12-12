package io

import (
	"fmt"
	"log"
	"strconv"
)

// ScanNumber - Request CLI input for number
func ScanNumber() float64 {
	var input string
	fmt.Scan(&input)

	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatalln("Incorrect number")
	}

	return float64(num)
}
