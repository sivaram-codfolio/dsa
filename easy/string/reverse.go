package main

import "fmt"

func reverse(content string) string {
	result := []rune(content)
	n := len(result)

	for i := 0; i < n/2; i++ {
		result[i], result[n-1-i] = result[n-1-i], result[i]
	}

	return string(result)
}

func main() {
	content := "abcd"
	fmt.Println("Original content :: ", content)
	fmt.Println("Reverse content :: ", reverse(content))
}
