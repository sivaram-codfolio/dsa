package main

import (
	"fmt"
	"strings"
)

func backtrack(s string, start int, path []string, result *[]string) {
	if len(path) == 4 && start == len(s) {
		*result = append(*result, strings.Join(path, "."))
		return
	}

	if len(path) == 4 {
		return
	}

	for length := 1; length <= 3; length++ {
		if start+length > len(s) {
			break
		}

		part := s[start : start+length]

		if (len(part) > 1 && part[0] == '0') || (len(part) == 3 && part > "255") {
			continue
		}

		backtrack(s, start+length, append(path, part), result)
	}
}

func restoreIpAddresses(s string) []string {
	result := []string{}
	backtrack(s, 0, []string{}, &result)
	return result
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
