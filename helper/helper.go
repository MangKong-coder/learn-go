package helper

import "strings"

func ValidateInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= int(remainingTickets)

	return isValidEmail, isValidName, isValidTickets
}
