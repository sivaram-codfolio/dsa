package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}
	parts := strings.Split(path, "/")

	for _, part := range parts {
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if part != "" && part != "." {
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/home/"))          // Output: "/home"
	fmt.Println(simplifyPath("/../"))            // Output: "/"
	fmt.Println(simplifyPath("/home//foo/"))     // Output: "/home/foo"
	fmt.Println(simplifyPath("/a/./b/../../c/")) // Output: "/c"
}
