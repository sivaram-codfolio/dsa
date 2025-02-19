package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		i--
		length++
	}

	return length
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))             // output: 5
	fmt.Println(lengthOfLastWord("luffy is still joyboy  ")) // output: 6
}
