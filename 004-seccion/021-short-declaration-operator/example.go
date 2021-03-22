package main

import "fmt"

func main() {
	x := 3
	y := "|Esteban Garviso"
	z := true
	a := 2*x + 3*x ^ 2 - 4
	fmt.Println(x)
	x = 33
	fmt.Println(x)
	fmt.Println(y)
	y = y + ` is the "best".|`
	fmt.Println(z)
	z = false
	fmt.Println(z)
	fmt.Println(x, y, z)
	fmt.Println(a)
}
