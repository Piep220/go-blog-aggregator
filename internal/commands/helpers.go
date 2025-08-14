package commands

import (
	"strings"
)

// cleanInput processes the input text by converting it to lowercase and splitting it into words.
// It returns a slice of strings containing the cleaned words.
func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}