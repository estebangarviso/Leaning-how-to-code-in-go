package main

import "fmt"

func main() {
	fmt.Println("My first go program")
	foo()
	fmt.Println("something else")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}
func foo() {
	fmt.Println("I'm in foo")
}
func bar() {
	fmt.Println("End")
}

//control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
