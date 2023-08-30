package main

import ("fmt"
		"strings"
	   )

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint =50

	fmt.Println("Welcome to our",conferenceName,"booking application")
	fmt.Printf("We have total %v tickets and %v tickets are still available\n",conferenceTickets,remainingTickets) // fomatted using the Printf() function 
	fmt.Println("Get your tickets here to attend")



	
	// datatypes in go
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// var bookings [50] string // array declartion in go 
	var bookings [] string // slice declaration in go 



	// taking user input 
	// fmt.Println("Enter the first name")
	// fmt.Scan(&firstName)
	// fmt.Println("Enter the last name")
	// fmt.Scan(&lastName)
	// fmt.Println("Enter the email")
	// fmt.Scan(&email)
	// fmt.Println("Enter the Number of user tickets")
	// fmt.Scan(&userTickets)


	// infinite for loop
	for{
	fmt.Println("Enter the first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter the last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter the email")
	fmt.Scan(&email)
	fmt.Println("Enter the Number of user tickets")
	fmt.Scan(&userTickets)
	// bookings = append(bookings,firstName+" "+lastName) 
	if userTickets>remainingTickets{
		fmt.Println("We only have ",remainingTickets, "tickets remaining ")
		continue
	}
	fmt.Println("These are all our bookings :",bookings)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets, email);


	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName+" "+lastName // inserting an element in an array 
	bookings = append(bookings,firstName+" "+lastName) // inserting an element in a slice 
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)
	
	var firstNames [] string
	// for each loop 
	for _, booking := range bookings{
		var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames,names[0])


	}
	fmt.Println("These are all our bookings :",firstNames)

	var noTicketsRemaining bool = remainingTickets == 0
	if  noTicketsRemaining{
		fmt.Println("Our conference is booked out. Come back next year")
		break
	}
	}




	// // arrays in go
	// // bookings [0] = "Rabin"
	// fmt.Printf("The Whole array :%v\n",bookings)
	// fmt.Printf("The first value :%v\n",bookings[0])
	// fmt.Printf("The array type :%T\n",bookings)
	// fmt.Printf("The length of array :%v\n",len(bookings))

	// slices in go 
	// fmt.Printf("The Whole slice :%v\n",bookings)
	// fmt.Printf("The first value :%v\n",bookings[0])
	// fmt.Printf("The slice type :%T\n",bookings)
	// fmt.Printf("The length of slice :%v\n",len(bookings))




	// type checking
	// fmt.Printf("conferenceName is of type %T\n",conferenceName)
	// fmt.Printf("remainingTickets is of type %T\n",conferenceTickets)
}