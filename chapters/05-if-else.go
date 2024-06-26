package main

import (
	"fmt"
	"strings"
)

func ifelse() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter your First Name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your E-mail address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			//book tickets in system
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive confirmation at %v email address!\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// print only first names
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out! Come back next year!!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name of last name is too short!!")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign!")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid Ticket Number!")
			}
		}

	}
	// Switch statement example
	city := "London"

	switch city {
	case "New York":
		// booking for New York conference
	case "Singapore", "Hong Kong":
		// booking for Singapore & Hong Kong conference
	case "London", "Berlin":
		// booking for London & Berlin conference
	case "Mexico City":
		// booking for Mexico City conference
	default:
		fmt.Print("No valid city selected")
	}

}
