package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) {
	var textReverse string
	s = strings.ToLower(s)

	for i := len(s) - 1; i >= 0; i-- {
		textReverse += string(s[i])
	}

	if strings.ToLower(s) == textReverse {
		fmt.Println("Is palyndrome")
	} else {
		fmt.Println("Is not palyndrome")
	}
}

func main() {
	slice := []string{"a", "b", "c", "d", "e"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	isPalindrome("Ama")
}
