package main

import "fmt"

func countCharacters(content string) map[rune]int {
	result := make(map[rune]int)
	for _, val := range content {
		result[val]++
	}

	return result
}

func main() {
	content := "abcd-abcd"
	fmt.Println("Count characters :: ", countCharacters(content))
}
