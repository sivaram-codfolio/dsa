package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phoneMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno",
		'7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	var result []string
	var backtrack func(index int, combination string)

	backtrack = func(index int, combination string) {
		if index == len(digits) {
			result = append(result, combination)
			return
		}

		letters := phoneMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, combination+string(letters[i]))
		}
	}

	backtrack(0, "")
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}
