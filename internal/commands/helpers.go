package commands

import (
	"fmt"
	"strings"
	"github.com/Piep220/go-blog-aggregator/internal/database"
)

// cleanInput processes the input text by converting it to lowercase and splitting it into words.
// It returns a slice of strings containing the cleaned words.
func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}