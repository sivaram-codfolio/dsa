package main

import "fmt"

func isPalindrome(content string) bool {
	list := []rune(content)
	n := len(list)

	for i := 0; i < n/2; i++ {
		if list[i] != list[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	content := "amma"
	fmt.Println("Is palindrome :: ", isPalindrome(content))
}
