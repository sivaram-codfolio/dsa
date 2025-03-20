package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordLen := len(words[0]) // Length of each word
	wordCount := make(map[string]int)

	// Build frequency map of words
	for _, word := range words {
		wordCount[word]++
	}

	var result []int

	// Try shifting starting position within word length
	for i := 0; i < wordLen; i++ {
		left, right := i, i
		currCount := make(map[string]int)
		count := 0 // Number of matched words

		// Sliding window
		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen

			// If word exists in words list
			if freq, found := wordCount[word]; found {
				currCount[word]++
				count++

				// If word appears too many times, shift window
				for currCount[word] > freq {
					leftWord := s[left : left+wordLen]
					currCount[leftWord]--
					left += wordLen
					count--
				}

				// If valid window, add start index
				if count == len(words) {
					result = append(result, left)
				}
			} else { // Reset window if invalid word found
				currCount = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return result
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words)) // Expected Output: [0, 9]
}
