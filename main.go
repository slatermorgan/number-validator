package main

import (
	"fmt"
	"strconv"
)

// weightingFactor is a lookup for what to multiply each digit by
var weightingFactor = [9]int{10, 9, 8, 7, 6, 5, 4, 3, 2}

const expectedNumberLength = 10

// divisor is used to divide the total by (11 for the mod11 algorithm)
const divisor = 11

// The check digit value if a result of 11 is found.
const elevenCheckDigit = 0

// IsValidNumber checks if a number passes the mod11 algorithm.
// arguement must be a 10 digit number
func IsValidNumber(input string) bool {
	if len(input) != expectedNumberLength {
		return false
	}

	var total int
	var lastDigit int
	for k, char := range input {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}

		// -1 as k is the index
		if k < expectedNumberLength-1 {
			total += digit * weightingFactor[k]
		}

		lastDigit = digit
	}

	remainder := total % divisor

	checkDigit := divisor - remainder

	// If the result is 11 then a check digit of 0 is used.
	if checkDigit == 11 {
		return lastDigit == elevenCheckDigit
	}

	return lastDigit == checkDigit
}

const printTemplate = "%s returns IsValidNumber = %t"

func main() {
	numbers := []string{
		"5990128088",
		"1275988113",
		"9876543210",
		"4536026665",
		"5990128087",
		"4536016660",
		"453601666",
		"qwertyuiop",
	}

	for _, num := range numbers {
		fmt.Println(fmt.Sprintf(printTemplate, num, IsValidNumber(num)))
	}
}
