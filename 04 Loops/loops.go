package main

import (
	"fmt"
)

func main() {
	// there is only for loop availabe in go but it can be used in three different ways

	// Way - 1: As an infinite loop

	count := 1
	for {

		if count > 10 { // I am placing this break statement here so that the loop doesnt run infinitely
			break
		}
		fmt.Println("Hello this is count", count)
		count++
	}

	// Way - 2: Loop with condition
	a := 1
	b := 5
	for a < b {
		fmt.Println("The value of a is", a)
		a++
	}

	// Way - 3 Using for loop as a for-each loop
	var num = []int{1, 99, 56, 0, 34, 9}
	for _, value := range num {
		fmt.Println("The value of num is", value)

	}
}
