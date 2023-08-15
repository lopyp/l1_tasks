package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid, 0
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0, 1
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	target := 6

	index, err := binarySearch(arr, target)

	if err == 1 {
		fmt.Printf("Element %d not found in array\n", target)
	} else {
		fmt.Printf("Element %d found at index %d in array\n", target, index)
	}
}
