package main

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := "1"

	for i := 2; i <= n; i++ {
		var sb strings.Builder
		count := 1
		for j := 1; j < len(prev); j++ {
			if prev[j] == prev[j-1] {
				count++
			} else {
				sb.WriteString(fmt.Sprintf("%d%c", count, prev[j-1]))
				count = 1
			}
		}
		sb.WriteString(fmt.Sprintf("%d%c", count, prev[len(prev)-1]))
		prev = sb.String()
	}

	return prev
}

func main() {
	fmt.Println(countAndSay(4)) // Output: "1211"
	fmt.Println(countAndSay(5)) // Output: "111221"
	fmt.Println(countAndSay(6)) // Output: "312211"
}
