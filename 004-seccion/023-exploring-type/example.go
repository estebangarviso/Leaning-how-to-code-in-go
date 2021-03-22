package main

import "fmt"

var y = 42

//	DECLARE the variable with the IDENTIFIER "z"
//	is of TYPE string
//	and I ASSIGN the value "Shaken, not stirred"
//	Should be omited type string from declaration of var "z";
//	it will be inferred from the right-hand sidego-lint
var z = "Shaken, not stirred"
var x = `James said, "Shaken, 

not stirred"`

//
func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	//	error: cannot use 43 (type int) as type string in assignment
	//	z = 43
	//	fmt.Println(z)
	//	fmt.Printf("%T\n", z)
	//	error: cannot use 43 (type int) as type string in assignment
	//	x = 43
	//	fmt.Println(x)
	//	fmt.Printf("%T\n", x)
}
