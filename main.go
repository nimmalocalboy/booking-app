package main

import (
	"fmt"
	"time"
)

const conferenceName = "GO conference"
const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	userEmail   string
	userTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, userEmail, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInputs(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, userEmail)

			go sendTicket(userTickets, firstName, lastName, userEmail)

			firstNames := getFirstNames()
			fmt.Printf("The first names are %v --> \n ", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First and last name should have a more than 2 letters")
			}
			if !isValidEmail {
				fmt.Println("Please check your entered email id - you are missing @ in the email")
			}
			if !isValidTicketNumber {
				fmt.Println("Please check the entered ticket number2")
			}
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application!", conferenceName)
	fmt.Printf("We have a total of %v tickets and the tickets available for purchase are : %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Please book your tickets here")
}

func getFirstNames() []string {
	fNames := []string{}
	for _, booking := range bookings {
		fNames = append(fNames, booking.firstName)
	}
	return fNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint
	//ask the user for name and no of tickets
	fmt.Println("Please enter your first name :")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name :")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email address :")
	fmt.Scan(&userEmail)
	fmt.Println("Please enter number of tickets you want to purchase :")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets
	//create a map of user data
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		userEmail:   userEmail,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("Hurry up!! Only %v tickets are remaining!!\n", remainingTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("Booking list --> %v \n", bookings)
}

func sendTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket :\n %v \n to email address %v\n", ticket, userEmail)
	fmt.Println("################")
}
