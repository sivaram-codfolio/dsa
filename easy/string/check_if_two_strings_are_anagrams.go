package main

import "fmt"

func isAnagrams(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	freq := make(map[rune]int)

	for _, ch := range str1 {
		freq[ch]++
	}

	for _, ch := range str2 {
		freq[ch]--
		if freq[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	str1 := "apple"
	str2 := "paple"

	fmt.Println("Are anagrams? ", isAnagrams(str1, str2))
}
