package main

import (
	"strings"
)

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	ValidName := len(firstName) >= 2 && len(lastName) >= 2
	ValidEmail := strings.Contains(email, "@")
	ValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return ValidName, ValidEmail, ValidTicketNumber
}
