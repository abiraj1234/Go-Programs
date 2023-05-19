package main

import (
	"fmt"
	"strings"
)

const bookingTickets int = 40

var BookMyshow = "revengers Phase 1"
var remainingTickets uint = 40
var bookings = []string{}

func main() {

	UserTicket(BookMyshow, bookingTickets, remainingTickets)

	// used infinte loops here
	for {
		firstName, lastName, email, userTickets := getUserinput()
		// vaildation function calls
		isVaildnames, isVaildEmails, isVaildUserickets := vaildateUserInput(firstName, lastName, email, userTickets)
		// check the tickets
		if isVaildnames && isVaildEmails && isVaildUserickets {

			bookTickets(userTickets, firstName, lastName, email)
			//calling the PrintNames func
			firstName := getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstName)

			// checking the remaining tickets
			if remainingTickets == 0 {
				fmt.Printf("Oopss!!Tickets Are Sold Out,.....")
				break
			}
		} else {

			if !isVaildnames {
				fmt.Println("Your First name or Last name are Invaild")
			}
			if !isVaildEmails {
				fmt.Println("Your email is Invaild")
			}
			if !isVaildUserickets {
				fmt.Println("your Enterd The Invaild Input")
			}
			break
		}

	}
}

func UserTicket(showName string, Tickets int, remaining uint) {
	fmt.Printf("Welcome to the %v  Bookingmyshow \n", BookMyshow)
	fmt.Printf("We Have total %v Tickets Are Available,  %v RemainingTickets.\n ", bookingTickets, remainingTickets)
	fmt.Print("Grab Your Tickets Here\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func vaildateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isVaildnames := len(firstName) >= 2 && len(lastName) >= 2
	isVaildEmails := strings.Contains(email, "@")
	isVaildUserickets := userTickets > 0 && userTickets <= remainingTickets

	return isVaildnames, isVaildEmails, isVaildUserickets
}

func getUserinput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask the username

	fmt.Println("Enter your firstName : ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your LastName: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter the Number of Tickets : ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// show  the names booked the tickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for %v booking Tickets. You will recive your confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v\n", remainingTickets, BookMyshow)
}
