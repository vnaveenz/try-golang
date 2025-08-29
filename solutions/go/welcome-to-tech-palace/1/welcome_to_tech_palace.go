package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	dsplyCustName := strings.ToUpper(customer)
	message := fmt.Sprintf("Welcome to the Tech Palace, %s", dsplyCustName)
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	message := fmt.Sprintf("%s\n%s\n%s",strings.Repeat("*",numStarsPerLine),welcomeMsg,strings.Repeat("*",numStarsPerLine))
	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(newMsg)
}
