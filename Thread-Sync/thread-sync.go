package main

import ("fmt"
		"sync"
		"time")
var wg = sync.WaitGroup{}
func main(){
	fmt.Println("This is an apple")
	wg.Add(1) // this function adds a number of threads which the main thread must wait for
	go greet()
	fmt.Println("This is a mango")
	wg.Wait() // waits until wait group i.e the threads that were added before counter is zero
}
 
func greet(){
	time.Sleep(10*time.Second)
	fmt.Println("This is the concurrent process running in the background")
	wg.Done() // Decrements the wait group counter by one 
}
	