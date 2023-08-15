package main

import (
	"fmt"
)

func multiplyByTwo(in <-chan int, out chan<- int) {
	for x := range in {
		out <- 2 * x
	}
	close(out)
}

func main() {
	data := []int{1, 2, 3}

	in := make(chan int, len(data))
	out := make(chan int, len(data))

	go multiplyByTwo(in, out)

	for _, v := range data {
		in <- v
	}
	close(in)

	for res := range out {
		fmt.Println(res)
	}
}
