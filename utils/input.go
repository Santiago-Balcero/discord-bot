package utils

import "strings"

func ClearString(input string) string {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	input = strings.TrimSpace(input)
	return input
}
