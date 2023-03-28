package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidNumber(t *testing.T) {
	tests := []struct {
		testDescription string
		number          string
		expect          bool
	}{
		{
			"Valid 1",
			"5990128088",
			true,
		},
		{
			"Valid 2",
			"1275988113",
			true,
		},
		{
			"Valid Ending with 0",
			"9876543210",
			true,
		},
		{
			"Valid 3",
			"4536026665",
			true,
		},
		{
			"Invalid 1",
			"5990128087",
			false,
		},
		{
			"Invalid 2",
			"4536016660",
			false,
		},
		{
			"Wrong Length",
			"453601666",
			false,
		},
		{
			"Wrong Characters",
			"qwertyuiop",
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.testDescription, func(t *testing.T) {
			isValid := IsValidNumber(test.number)

			assert.Equal(t, test.expect, isValid)
		})
	}

}
