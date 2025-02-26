# Data Structures and Algorithms (DSA) in Go

A collection of data structures and algorithms implemented in Go. This repository serves as a reference for developers looking to understand and apply DSA concepts effectively in Go projects.

🚀 *DSA Covered**

🏗 **Sort - Bubble**

Given an integer array nums sorted in non-decreasing order.

📌 *Example Execution*

```bash
go run cmd/sort/bubble.go
```

📜 *Sample Output*

```bash
[1 1 2 3 4 5 6 9]
```

🏗 **Quick - Bubble**

Given an integer array nums sorted in non-decreasing order.

📌 *Example Execution*

```bash
go run cmd/sort/quick.go
```

📜 *Sample Output*

```bash
[1 1 2 3 4 5 6 9]
```

🏗 **Duplicate Sorted Array**

Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

📌 *Example Execution*

```bash
go run cmd/duplicate_sorted_array.go
```

📜 *Sample Output*

```bash
Output: 2, nums1 = [1 2]
Output: 8, nums2 = [0 1 2 3 4 5 6 7]
```

🏗 **Remove Element**

Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

📌 *Example Execution*

```bash
go run cmd/remove_element.go
```

📜 *Sample Output*

```bash
Input: nums = [3 2 2 3], val = 3
Output: 2, nums = [2 2]

Input: nums = [0 1 2 2 3 0 4 2], val = 2
Output: 5, nums = [0 1 3 0 4]
```

🏗 **Search Insert Position**

The Search Insert Position problem is about finding the index at which a target value should be inserted into a sorted array, such that the array remains sorted.

📌 *Example Execution*

```bash
go run cmd/search_insert_position.go
```

📜 *Sample Output*

```bash
2
1
4
0
```

🏗 **Length of Last Word**

Given a string s consisting of words and spaces, return the length of the last word in the string. A word is a maximal substring consisting of non-space characters only.

📌 *Example Execution*

```bash
go run cmd/length_last_word.go
```

📜 *Sample Output*

```bash
5
6
```

🏗 **Plus One**

You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

📌 *Example Execution*

```bash
go run cmd/plus_one.go
```

📜 *Sample Output*

```bash
[1 3 0]
[4 3 2 2]
[1 0]
```

🏗 **Add Binary**

Given two binary strings a and b, return their sum as a binary string.

📌 *Example Execution*

```bash
go run cmd/add_binary.go
```

📜 *Sample Output*

```bash
100
10101
0
1000
```

🏗 **Sqrt(x)**

Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

📌 *Example Execution*

```bash
go run cmd/sqrt.go
```

📜 *Sample Output*

```bash
2
2
4
1
0
```

🏗 **Longest Palindromic String**

Given a string s, return the longest `palindromic substring` in s.

📌 *Example Execution*

```bash
go run cmd/longest_palindromic_string.go
```

📜 *Sample Output*

```bash
Longest Palindromic Substring: bab
```

🏗 **Median of Two Sorted Array**

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

📌 *Example Execution*

```bash
go run cmd/median_of_two_sorted_array.go
```

📜 *Sample Output*

```bash
Median: 2
Median: 2.5
```

🏗 **Reverse Integer**

Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

📌 *Example Execution*

```bash
go run cmd/reverse_integer.go
```

📜 *Sample Output*

```bash
321
-321
21
0
```

🏗 **3Sum**

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

📌 *Example Execution*

```bash
go run cmd/3sum.go
```

📜 *Sample Output*

```bash
[[-1 -1 2] [-1 0 1]]
[]
[[0 0 0]]
```

🤝 **Contribution**

We welcome contributions! If you’d like to enhance the repository, you can:
-	📝 Add new data structures or algorithms
-	⚡ Optimize existing implementations
-	📖 Improve documentation

Steps to Contribute
1.	Fork the repository
2.	Create a new branch for your changes
3.	Commit your modifications
4.	Open a pull request

Your contributions help improve this project for everyone!

📜 **License**

This project is licensed under the MIT License. See the LICENSE file for details.
