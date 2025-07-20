package main

import "fmt"

func findNonRepeatingChar(str string) rune {
	freq := make(map[rune]int)

	for _, val := range str {
		freq[val]++
	}
	fmt.Println("freq :: ", freq)

	for _, val := range str {
		if freq[val] == 1 {
			return val
		}
	}

	return 0
}

func main() {
	str := "swiss"
	ch := findNonRepeatingChar(str)
	fmt.Printf("Find a non repeating character :: %c\n", ch)
}
