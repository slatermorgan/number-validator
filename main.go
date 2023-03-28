package main

import (
	"strconv"
)

var weightingFactor = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

const checkDigitIndex = 9

const divisor = 11

func IsValidNumber(numString string) bool {
	if len(numString) != 10 {
		return false
	}

	var total int
	var lastDigit int
	for k, char := range numString {
		digit := string(char)
		num, err := strconv.Atoi(digit)
		if err != nil {
			return false
		}

		if k == checkDigitIndex {
			lastDigit = num

			break
		}

		total += num * weightingFactor[k]
	}

	remainder := total % divisor

	checkDigit := divisor - remainder

	return lastDigit == checkDigit
}
