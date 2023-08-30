package main

import "fmt"

func main() {
	var conferenceName = "Go Conference" // declaration with initilization
	const conferenceTickets = 50         // a constant variable
	var name string                      // declaration without initilization
	name = "Rabin"
	hobby := "cricket" // this is another way of declaring a variable in go

	fmt.Println("The conference name is ", conferenceName)
	fmt.Println("There are", conferenceTickets, "conference tickets")
	fmt.Println(name, "has already booked the conference tickets")
	fmt.Println("My hobby is", hobby)

}
