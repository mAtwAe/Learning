package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are stil available.")

	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Println(conferenceN ame)

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// isValidCity := city != "Singapore" && city != "London"

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName)

			firstNames := getFirstName()

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

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number ticket is invalid")
			}

			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	//execute code for new york
	// case "Singapore":
	// 	//execute code for singaport
	// case "London", "Berlin":
	// 	//execute code for London and Berlin
	// case "Mexico City":
	// 	//execute code for Mexico City
	// case "Hong Kong":
	// 	//execute code for Hong Kong
	// default:
	// 	fmt.Println("No valid city selected")
	// }

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are stil available.")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
	fmt.Printf("The whole slice: %v\n", bookings)
}
