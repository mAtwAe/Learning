package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are stil available.")

	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Println(conferenceN ame)

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// isValidCity := city != "Singapore" && city != "London"

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(remainingTickets, userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstName()

		fmt.Printf("The first names of bookings are: %v \n", firstNames)

		// fmt

		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("S lice length: %v\n", len(bookings))

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are stil available.")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create map for user
	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
	fmt.Printf("The whole slice: %v\n", bookings)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}
