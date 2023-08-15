package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{} = 10
	var s interface{} = "hello"
	var b interface{} = true
	var c interface{} = make(chan int)

	fmt.Println("Type of i:", reflect.TypeOf(i))
	fmt.Println("Type of s:", reflect.TypeOf(s))
	fmt.Println("Type of b:", reflect.TypeOf(b))
	fmt.Println("Type of c:", reflect.TypeOf(c))
}
