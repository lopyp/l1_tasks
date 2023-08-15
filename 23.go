package main

import "fmt"

func removeElement(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func main() {
	s := []int{1, 3, 4, 5}
	fmt.Println("before:", s)
	s = removeElement(s, 2)
	fmt.Println("after:", s)
}
