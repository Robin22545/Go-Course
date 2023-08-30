package main
import ("fmt"
		"time")

func main(){
	var name string
	fmt.Println("This is a cat")
	go greet() // the go keyword is used to create a concurrent process 
	fmt.Println("This is an apple")
	for{
		fmt.Println("Enter your name")
		fmt.Scan(&name)
	}
}

func greet(){
	time.Sleep(10*time.Second) // setting a delay of 10 sec
	fmt.Println("Hello I am the concurrent process which was running in backgound")
}