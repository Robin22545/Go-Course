// Taking user input in go
package main

import "fmt"

func main() {

	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name) // we pass the address of the variable
	fmt.Println("Your name is", name)
}
