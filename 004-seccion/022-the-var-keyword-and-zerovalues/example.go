package main

import "fmt"

// the VARIABLE with the IDENTIFIER "y"
// is a global variable of the program
var y = 3

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z" = 0
var z int

// DECLARE there is a VARIABLE with the IDENTIFIER "a"
// and that the VARIABLE with the IDENTIFIER "a" is of TYPE bool
// ASSIGNS the ZERO VALUE of TYPE bool to "a" = false
var a bool

// DECLARE there is a VARIABLE with the IDENTIFIER "b"
// and that the VARIABLE with the IDENTIFIER "b" is of TYPE string
// ASSIGNS the ZERO VALUE of TYPE string to "b" = ""
var b string

// there more ZERO value like float = 0.0 and
// pointers, funtions, interfaces, slices, channels and maps = nil
func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 6
	// x is a local variable of the function
	fmt.Println(x)
	fmt.Println("Printing in main function: ", y, "\nPrinting in main function: ", z)
	foo()
}
func foo() {
	fmt.Println("Printing again in foo function: ", y, "\nPrinting again in foo function: ", z)
	// if I put the next statement:
	// fmt.Println(x)
	// is not going to work, because is a LOCAL VARIABLE of main() no of foo()
	// error: undefined: x
	fmt.Println("Printing ZERO VALUES of int, bool and string:", z, a, b)
}
