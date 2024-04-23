package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClearString(t *testing.T) {
	result := ClearString("  CALYpsa la Loca ")
	assert.Equal(
		t,
		"calypsalaloca",
		result,
		"Should clear upper case letters and spaces",
	)
}

func TestClearStringOnlySpaces(t *testing.T) {
	result := ClearString("      ")
	assert.Equal(
		t,
		"",
		result,
		"Should return empty string when handling string of empty spaces",
	)
}

func TestClearStringEmpty(t *testing.T) {
	result := ClearString("")
	assert.Equal(
		t,
		"",
		result,
		"Should return empty string",
	)
}
