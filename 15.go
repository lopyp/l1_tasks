package main

import "fmt"

var justString string

// паника если строка меньше 100 байт
func someFunc() {
	v := "somestring"
	a := len(v)
	justString = v[:a]
	fmt.Print(justString)
}

func main() {
	someFunc()
}
