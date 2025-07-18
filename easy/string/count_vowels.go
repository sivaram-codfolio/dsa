package main

import "fmt"

func countVowels(content string) int {
	vowel := "aeiouAEIOU"
	count := 0

	for _, ch := range content {
		for _, v := range vowel {
			if ch == v {
				count++
			}
		}
	}

	return count
}

func main() {
	content := "Golang is great"
	fmt.Println("Vowels count :: ", countVowels(content))
}
