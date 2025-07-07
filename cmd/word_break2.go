package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	memo := make(map[int][]string)
	return dfs(s, 0, wordSet, memo)
}

func dfs(s string, start int, wordSet map[string]bool, memo map[int][]string) []string {
	if val, ok := memo[start]; ok {
		return val
	}

	if start == len(s) {
		return []string{""} // Return list with empty string to build up
	}

	var result []string
	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]
		if wordSet[word] {
			sublist := dfs(s, end, wordSet, memo)
			for _, sub := range sublist {
				if sub == "" {
					result = append(result, word)
				} else {
					result = append(result, word+" "+sub)
				}
			}
		}
	}

	memo[start] = result
	return result
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	result := wordBreak(s, wordDict)
	fmt.Println("All possible sentences:")
	for _, sentence := range result {
		fmt.Println(sentence)
	}
}
