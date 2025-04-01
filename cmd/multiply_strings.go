package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	// Edge case: If either number is "0", return "0"
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	result := make([]int, m+n) // Maximum length of product

	// Multiply each digit of num1 and num2
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1, p2 := i+j, i+j+1
			sum := mul + result[p2]

			result[p2] = sum % 10  // Store ones place
			result[p1] += sum / 10 // Carry to next position
		}
	}

	// Convert result array to string (skip leading zeros)
	var res string
	for _, num := range result {
		if !(num == 0 && len(res) == 0) { // Skip leading zeros
			res += string(num + '0')
		}
	}

	return res
}

func main() {
	fmt.Println(multiply("123", "456")) // Output: "56088"
	fmt.Println(multiply("2", "3"))     // Output: "6"
	fmt.Println(multiply("0", "1234"))  // Output: "0"
}
