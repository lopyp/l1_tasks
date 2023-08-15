package main

import (
	"fmt"
)

func intersection(set1, set2 []int) []int {
	m := make(map[int]bool)
	intersect := []int{}

	for _, item := range set1 {
		m[item] = true
	}

	for _, item := range set2 {
		if _, ok := m[item]; ok {
			intersect = append(intersect, item)
		}
	}

	return intersect
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	intersect := intersection(set1, set2)

	fmt.Println(intersect) // [4 5]
}
