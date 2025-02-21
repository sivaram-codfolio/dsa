package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	maxLen := la
	if lb > la {
		maxLen = lb
	}

	result := make([]byte, 0, maxLen+1)
	carry := 0
	i, j := la-1, lb-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = append(result, byte(sum%2)+'0')
		carry = sum / 2
	}

	for k, n := 0, len(result); k < n/2; k++ {
		result[k], result[n-1-k] = result[n-1-k], result[k]
	}

	return string(result)
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("111", "1"))
}
