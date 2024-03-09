package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("test")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are stil available.")

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println(conferenceName)

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// firstName = "Ahmad"
		// lastName = "Albab"

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// bookings[0] = firstName + " " + lastName
		// bookings[10] = "Abu"

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Welcome to %v booking application\n", conferenceName)
			fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
			fmt.Printf("Get your tickets here to attend\n")
			fmt.Printf("The whole slice: %v\n", bookings)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			// fmt

			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("S lice length: %v\n", len(bookings))

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

}
