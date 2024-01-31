package helper

import (
	"strings"
)

// capitalize func name to export function to othe packages
func ValidateUserInput(fName string, lName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets
	//isValidCity := city == "Singapore" || city == "London"
	//!isValidCity
	return isValidName, isValidEmail, isValidTicketNum
}
