package main 
import ("fmt"
	   )

func main(){
	var ar =[] int {2,3,4,6}
	// Function call
	greet() 
	num(23)
	arr(ar)
	fmt.Println("The sum of 3 and 5 is",add(3,4))
	fmt.Println("Returning a function from an array",A())
	a,b,c:= B()
	fmt.Println(a,b,c)
	
}


// a simple function
func greet(){
fmt.Println("Hello Rabin")
}

// function which takes parameters
func num( val int){
	fmt.Println("The passed value is",val)
}

// passing array to a function
func arr(ar []int){
	fmt.Println("The content of the array is ",ar);
}

//  returning values from a function
func add(num1 int, num2 int) int {
	return num1 + num2
}

// returning an array from a function
func A() []int{
	var a  = [] int{1,2,3}
	return a
} 

// returning multiple values from a function

func B()(int, string,bool){
	return 23,"apple",false
}