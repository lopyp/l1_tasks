package main

import (
	"fmt"
)

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left, equal, right := []int{}, []int{}, []int{}

	for _, num := range arr {
		switch {
		case num < pivot:
			left = append(left, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			right = append(right, num)
		}
	}

	return append(append(quicksort(left), equal...), quicksort(right)...)
}

func main() {
	arr := []int{5, 2, 8, 3, 1, 9, 4, 6, 7}
	fmt.Println("Unsorted:", arr)
	fmt.Println("Sorted:", quicksort(arr))
}
