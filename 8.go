package main

import (
	"fmt"
)

func setBit(n int64, pos uint, bit bool) int64 {
	if bit {
		n |= (1 << pos)
	} else {
		n &^= (1 << pos)
	}
	return n
}

func main() {
	var n int64 = 2340

	fmt.Printf("%b\n", n)
	// Установить бит 1
	n = setBit(n, 1, true)
	fmt.Printf("%b\n", n)

	// Установить бит 0
	n = setBit(n, 2, false)
	fmt.Printf("%b\n", n)
}
