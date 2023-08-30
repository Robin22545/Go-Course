package main
import "fmt"

type data struct{

	Name string 
	RollNo string
	Year string 
	Sem  string
}

var values = make([]data,0) // an slice which can store struct instance as items
func main(){


// assigining values to the instance of the struct
var s1 = data{
	Name:"Rabin",
	RollNo:"20CSEC04",
	Year:"Final",
	Sem:"7th",
}

// we can also assign values to the struct as 
var s2 = data{} // creating the instance of the structure 
s2.Name = "Rabin Sharma" // we can assign values using the dot operator
fmt.Println(s2.Name)

values = append(values,s1) // storing the instance of the structure as element of the slice
fmt.Println("The first element of the array is ",values[0])

}