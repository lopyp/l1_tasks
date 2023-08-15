package main

import (
	"fmt"
	"unicode"
)

func isUnique(s string) bool {
	m := make(map[rune]bool)
	for _, r := range s {
		if unicode.IsLetter(r) {
			if m[unicode.ToLower(r)] {
				return false
			}
			m[unicode.ToLower(r)] = true
		}
	}
	return true
}

func main() {
	s1 := "asdfgA"
	s2 := "xcvbnclX"
	fmt.Println(s1, "is unique:", isUnique(s1))
	fmt.Println(s2, "is unique:", isUnique(s2))
}
