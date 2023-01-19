package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainigTickets = 50

	// fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainigTickets)
	fmt.Println("Get your tickets here to attend.")

}
