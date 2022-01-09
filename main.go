package main

import (
	"fmt"
	"learn-go/helper"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidEmail, isValidName, isValidTickets := helper.ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(uint(userTickets), firstName, lastName, email)

		fmt.Printf("These are the users: %v \n", bookings)

		firstNames := getFirstNames()

		fmt.Printf("The users: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("Name is too short")
		}
		if !isValidEmail {
			fmt.Println("Email is invalid")
		}
		if !isValidTickets {
			fmt.Println("Number of tickets is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome %v to our conference \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(userTickets),
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v ticekts. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v remaining tickets\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket \n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
