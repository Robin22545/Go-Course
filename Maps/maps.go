package main

import "fmt"

func main(){

	var data = make (map[string]string) // this is how we declare a map where both the keys and  values are string
	data["Name"] = "Rabin Sharma"
	data["RollNo"]="20CSEC04"
	data["Year"] = "Final"
	data["Sem"] = "7th"
	fmt.Println("The name of the customer is",data["Name"])
	fmt.Println("The entire data: ",data)
	// Iterating through a map uisng a loop
	fmt.Println("Now lets iterate through the map")
	for _,value := range data{
		fmt.Println(value)
	}


}