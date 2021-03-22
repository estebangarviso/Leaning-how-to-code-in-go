package main

import "fmt"

type nproduct int

var x nproduct

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3
	fmt.Println(x)
}
