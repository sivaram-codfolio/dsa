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

🏗 **Sort - Merge Array's**

Given two integer array nums merged and sorted in non-decreasing order.

📌 *Example Execution*

```bash
go run cmd/sort/merge_sort_arrays.go
```

📜 *Sample Output*

```bash
[1 2 3 4 5 8 9]
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

🏗 **Container With Most Water**

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

📌 *Example Execution*

```bash
go run cmd/container_with_most_water.go
```

📜 *Sample Output*

```bash
49
1
```

🏗 **String to Integer (atoi)**

Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.

The algorithm for myAtoi(string s) is as follows:

Whitespace: Ignore any leading whitespace (" ").
Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
Return the integer as the final result.

📌 *Example Execution*

```bash
go run cmd/string_to_integer.go
```

📜 *Sample Output*

```bash
42
-42
4193
0
-2147483648
```

🏗 **Integer to Roman**

Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:

If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.
Given an integer, convert it to a Roman numeral.

📌 *Example Execution*

```bash
go run cmd/integer_to_roman.go
```

📜 *Sample Output*

```bash
III
LVIII
MCMXCIV
```

🏗 **Generate Parentheses**

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

📌 *Example Execution*

```bash
go run cmd/generate_parentheses.go
```

📜 *Sample Output*

```bash
[((())) (()()) (())() ()(()) ()()()]
[()]
```

🏗 **Letter Combinations of a Phone Number**

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

📌 *Example Execution*

```bash
go run cmd/letter_combination_of_phone_number.go
```

📜 *Sample Output*

```bash
[ad ae af bd be bf cd ce cf]
[]
[a b c]
```

🏗 **Divide Two Integers**

Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.

📌 *Example Execution*

```bash
go run cmd/divide_two_integers.go
```

📜 *Sample Output*

```bash
3
-2
0
1
2147483647
```

🏗 **Majority Element**

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

📌 *Example Execution*

```bash
go run cmd/majority_element.go
```

📜 *Sample Output*

```bash
3
2
```

🏗 **Valid Sudoku**

Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

📌 *Example Execution*

```bash
go run cmd/valid_sudoku.go 
```

📜 *Sample Output*

```bash
true
```

🏗 **Intersection of Two Linked Lists**

Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

📌 *Example Execution*

```bash
go run cmd/intersection_of_two_linked_lists.go
```

📜 *Sample Output*

```bash
List A:
4 -> 1 -> 8 -> 4 -> 5 -> nil
List B:
5 -> 6 -> 1 -> 8 -> 4 -> 5 -> nil
Intersection at node with value: 8
```

🏗 **Climbing Stairs**

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

📌 *Example Execution*

```bash
go run cmd/climbing_stairs.go
```

📜 *Sample Output*

```bash
3
8
89
```

🏗 **Remove Duplicates from Sorted List**

Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

📌 *Example Execution*

```bash
go run cmd/remove_duplicate_from_sorted_list.go
```

📜 *Sample Output*

```bash
Original List:
1 → 1 → 2 → 3 → 3 → nil
List after removing duplicates:
1 → 2 → 3 → nil
```

🏗 **Same Tree**

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

📌 *Example Execution*

```bash
go run cmd/same_tree.go
```

📜 *Sample Output*

```bash
Are the two trees identical? true
```

🏗 **Simplify Path**

You are given an absolute path for a Unix-style file system, which always begins with a slash '/'. Your task is to transform this absolute path into its simplified canonical path.

📌 *Example Execution*

```bash
go run cmd/simplify_path.go
```

📜 *Sample Output*

```bash
/home
/
/home/foo
/c
```

🏗 **Swap Nodes in Pairs**

Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

📌 *Example Execution*

```bash
go run cmd/simplify_path.go
```

📜 *Sample Output*

```bash
/home
/
/home/foo
/c
```

🏗 **Regular Expression Matching**

Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

📌 *Example Execution*

```bash
go run cmd/regular_expression_matching.go 
```

📜 *Sample Output*

```bash
true
true
false
```

🏗 **3Sum Closest**

Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.

📌 *Example Execution*

```bash
go run cmd/3sum_closest.go
```

📜 *Sample Output*

```bash
2
```

🏗 **Merge k Sorted Lists**

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

📌 *Example Execution*

```bash
go run cmd/merge_k_sorted_lists.go
```

📜 *Sample Output*

```bash
1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6 -> nil
```

🏗 **Remove Nth Node From End of List**

Given the head of a linked list, remove the nth node from the end of the list and return its head.

📌 *Example Execution*

```bash
go run cmd/remove_nth_node_from_end_of_list.go
```

📜 *Sample Output*

```bash
Original List: 1 → 2 → 3 → 4 → 5 → nil
After Removal: 1 → 2 → 3 → 5 → nil
```

🏗 **Reverse Nodes in k-Group**

Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

📌 *Example Execution*

```bash
go run cmd/reverse_nodes_k_group.go
```

📜 *Sample Output*

```bash
Original List: 1 → 2 → 3 → 4 → 5 → nil
After k-Reverse: 2 → 1 → 4 → 3 → 5 → nil
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
