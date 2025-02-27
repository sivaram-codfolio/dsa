package main

import "fmt"

func myAtoi(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	i := 0
	for i < n && s[i] == ' ' {
		i++
	}

	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	num := 0
	INT_MAX := 2147483647
	INT_MIN := -2147483648

	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if num > (INT_MAX-digit)/10 {
			if sign == 1 {
				return INT_MAX
			}
			return INT_MIN
		}

		num = num*10 + digit
		i++
	}

	return num * sign
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
}
