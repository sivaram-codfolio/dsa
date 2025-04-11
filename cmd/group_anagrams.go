package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		sorted := sortString(word)
		anagramMap[sorted] = append(anagramMap[sorted], word)
	}

	// Build result from map
	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

// Helper function to sort string characters
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
