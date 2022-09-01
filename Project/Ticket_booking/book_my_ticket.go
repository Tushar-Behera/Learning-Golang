package main

import "fmt"

func main() {

	// PREDEFINE VAR AS PER REQUIREMENT

	appName := "Book My Ticket 🙂"
	const totalTicket int = 50
	var remainingTicket uint = 50

	fmt.Printf("\nWelcome to %v\n", appName)
	fmt.Printf("Todays Ticket count is : %v and Remaining : %v Tickets are available\n", totalTicket, remainingTicket)
	fmt.Println("Book your ticket 🎫 now to enjoy your Shows 🤓 ")

	// INPUT FROM USER

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("\n\nPlease Enter your First Name : ")
	fmt.Scan(&firstName)

	fmt.Println("Please Enter your Last Name : ")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter your Email : ")
	fmt.Scan(&email)

	fmt.Println("No of Tickets to book :")
	fmt.Scan(&userTicket)

	remainingTicket = remainingTicket - userTicket

	fmt.Printf("\n\nHello Dear %v %v Thank You For Booking %v Tickets 🤩", firstName, lastName, userTicket)
	fmt.Printf("\nYour Conformation will be sended on %v this email . Enjoy 👍\n", email)

	fmt.Printf("Reamining %v Tickets Available ...Hurry Now ✨", remainingTicket)
}
