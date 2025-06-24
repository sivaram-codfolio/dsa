package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	var filtered []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	i, j := 0, len(filtered)-1
	for i < j {
		if filtered[i] != filtered[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	input := "A man, a plan, a canal: Panama"
	fmt.Println("Is palindrome?", isPalindrome(input)) // Output: true
}
