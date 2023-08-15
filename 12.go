package main

import (
	"fmt"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, item := range sequence {
		set[item] = true
	}

	for item := range set {
		fmt.Println(item)
	}
}
