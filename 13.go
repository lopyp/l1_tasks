package main

import "fmt"

func main() {
	a := 5
	b := 15
	fmt.Println("a =", a, "b =", b)

	a += b
	b = a - b
	a -= b
	fmt.Println("a =", a, "b =", b)
}
