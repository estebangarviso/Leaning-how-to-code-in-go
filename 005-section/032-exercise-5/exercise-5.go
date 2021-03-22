package main

import "fmt"

type nproduct int

var x nproduct
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T", y)
}
