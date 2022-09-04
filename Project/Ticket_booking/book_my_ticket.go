package main

import "fmt"

func main() {

	// PREDEFINE VAR AS PER REQUIREMENT

	appName := "Book My Ticket 🙂"
	const totalTicket int = 50
	var remainingTicket uint = 50
	bookings := []string{}
	fmt.Printf("\nWelcome to %v\n", appName)
	fmt.Printf("Todays Ticket count is : %v and Remaining : %v Tickets are available\n", totalTicket, remainingTicket)
	fmt.Println("Book your ticket 🎫 now to enjoy your Shows 🤓 ")

	// Use of For loop

	for {
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Printf("\n\nPlease Enter your First Name :\n")
		fmt.Scan(&firstName)

		fmt.Println("Please Enter your Last Name :")
		fmt.Scan(&lastName)

		fmt.Println("Please Enter your Email :")
		fmt.Scan(&email)

		fmt.Println("No of Tickets to book :")
		fmt.Scan(&userTicket)

		remainingTicket = remainingTicket - userTicket
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("\n\nHello Dear [ %v %v ] Thank You For Booking [ %v ] Tickets 🤩", firstName, lastName, userTicket)

		fmt.Printf("\nYour Confirmation will be sent to [ %v ] this email. Enjoy👍\n", email)

		fmt.Printf("\nHere are all the Booking Lists : %v %v 🎫\n", bookings, userTicket)

		fmt.Printf("\nReamining %v Tickets Available ...Hurry Now ✨\n", remainingTicket)
	}

}
