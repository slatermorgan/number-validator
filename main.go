package main

import (
	"strconv"
)

// weightingFactor is a lookup for what to multiply each digit by
var weightingFactor = [9]int{10, 9, 8, 7, 6, 5, 4, 3, 2}

// checkDigitIndex is the index of the check digit in the number string
const checkDigitIndex = 9

// divisor is used to divide the total by (11 for the mod11 algorithm)
const divisor = 11

// IsValidNumber checks if a number passes the mod11 algorithm.
// arguement must be a 10 digit number
func IsValidNumber(numString string) bool {
	if len(numString) != 10 {
		return false
	}

	var total int
	var lastDigit int
	for k, char := range numString {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}

		if k == checkDigitIndex {
			lastDigit = digit

			break
		}

		total += digit * weightingFactor[k]
	}

	remainder := total % divisor

	checkDigit := divisor - remainder

	return lastDigit == checkDigit
}
