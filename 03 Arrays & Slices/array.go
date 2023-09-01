// lets see how we can create static and dynamic arrays in go
// A slice is nothing but a dynamic array

package main

import "fmt"

func main() {
	var names [50]string // array declaration
	names[0] = "Rabin"   // inserting item into an array

	var fruits = [50]string{"Apple", "mango"} // array declaration with initilization
	fmt.Println(fruits[0], "is a healthy fruit")

	var schools []string               // declaring a slice (dynamic array)
	schools = append(schools, "ALSSS") // we can also insert item to an array or slice using the append function

}
