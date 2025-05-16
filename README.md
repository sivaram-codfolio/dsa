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

🏗 **4Sum**

Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

📌 *Example Execution*

```bash
go run cmd/4sum.go
```

📜 *Sample Output*

```bash
[[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
```

🏗 **Substring with Concatenation of All Words**

You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.

📌 *Example Execution*

```bash
go run cmd/substring_with_concatenation_of_all_words.go
```

📜 *Sample Output*

```bash
[0 9]
```

🏗 **Binary Tree Inorder Traversal**

Given the root of a binary tree, return the inorder traversal of its nodes' values.

📌 *Example Execution*

```bash
go run cmd/binary_tree_inorder_traversal.go
```

📜 *Sample Output*

```bash
[1 3 2]
```

🏗 **Merge Sorted Array**

You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

📌 *Example Execution*

```bash
go run cmd/merge_sorted_array.go
```

📜 *Sample Output*

```bash
[1 2 2 3 5 6]
```

🏗 **Maximum Depth of Binary Tree**

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

📌 *Example Execution*

```bash
go run cmd/maximum_depth_of_binary_tree.go
```

📜 *Sample Output*

```bash
3
```

🏗 **Pascal's Triangle**

Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

📌 *Example Execution*

```bash
go run cmd/pascal_triangle.go
```

📜 *Sample Output*

```bash
[[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
```

🏗 **Find First and Last Position of Element in Sorted Array**

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

📌 *Example Execution*

```bash
go run cmd/find_first_and_last_position_element_in_sorted_array.go
```

📜 *Sample Output*

```bash
[3 4]
```

🏗 **Restore IP Addresses**

A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s. You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.

📌 *Example Execution*

```bash
go run cmd/restore_ip_address.go
```

📜 *Sample Output*

```bash
[255.255.11.135 255.255.111.35]
```

🏗 **Next Permutation**

A permutation of an array of integers is an arrangement of its members into a sequence or linear order.

For example, for arr = [1,2,3], the following are all the permutations of arr: [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].
The next permutation of an array of integers is the next lexicographically greater permutation of its integer. More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, then the next permutation of that array is the permutation that follows it in the sorted container. If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

For example, the next permutation of arr = [1,2,3] is [1,3,2].
Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
Given an array of integers nums, find the next permutation of nums.

The replacement must be in place and use only constant extra memory.

📌 *Example Execution*

```bash
go run cmd/next_permutation.go
```

📜 *Sample Output*

```bash
[1 3 2]
[1 2 3]
[1 5 1]
```

🏗 **Longest Valid Parentheses**

Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses substring.

📌 *Example Execution*

```bash
go run cmd/longest_valid_parentheses.go
```

📜 *Sample Output*

```bash
2
4
0
```

🏗 **Count and Say**

The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the run-length encoding of countAndSay(n - 1).
Run-length encoding (RLE) is a string compression method that works by replacing consecutive identical characters (repeated 2 or more times) with the concatenation of the character and the number marking the count of the characters (length of the run). For example, to compress the string "3322251" we replace "33" with "23", replace "222" with "32", replace "5" with "15" and replace "1" with "11". Thus the compressed string becomes "23321511".

Given a positive integer n, return the nth element of the count-and-say sequence.

📌 *Example Execution*

```bash
go run cmd/count_and_say.go 
```

📜 *Sample Output*

```bash
1211
111221
312211
```

🏗 **Combination Sum**

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

📌 *Example Execution*

```bash
go run cmd/combination_sum.go 
```

📜 *Sample Output*

```bash
[[2 2 3] [7]]
[[2 2 2 2] [2 3 3] [3 5]]
```

🏗 **Combination Sum II**

Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

📌 *Example Execution*

```bash
go run cmd/combination_sum2.go
```

📜 *Sample Output*

```bash
[[1 1 6] [1 2 5] [1 7] [2 6]]
[[1 2 2] [5]]
```

🏗 **Multiply Strings**

Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

📌 *Example Execution*

```bash
go run cmd/multiply_strings.go
```

📜 *Sample Output*

```bash
56088
6
0
```

🏗 **Jump Game II**

You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

📌 *Example Execution*

```bash
go run cmd/jump_game2.go
```

📜 *Sample Output*

```bash
2
2
```

🏗 **First Missing Positive**

Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

📌 *Example Execution*

```bash
go run cmd/find_missing_positive.go
```

📜 *Sample Output*

```bash
3
2
1
```

🏗 **Permutations**

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

📌 *Example Execution*

```bash
go run cmd/permutations.go
```

📜 *Sample Output*

```bash
[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
```

🏗 **Permutations II**

Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

📌 *Example Execution*

```bash
go run cmd/permutations2.go
```

📜 *Sample Output*

```bash
[[1 1 2] [1 2 1] [2 1 1]]
```

🏗 **Rotate Image**

You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

📌 *Example Execution*

```bash
go run cmd/rotate_image.go
```

📜 *Sample Output*

```bash
[[7 4 1] [8 5 2] [9 6 3]]
```

🏗 **Group Anagrams**

Given an array of strings strs, group the anagrams together. You can return the answer in any order.

📌 *Example Execution*

```bash
go run cmd/group_anagrams.go
```

📜 *Sample Output*

```bash
[[eat tea ate] [tan nat] [bat]]
```

🏗 **Pow(x, n)**

Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

📌 *Example Execution*

```bash
go run cmd/pow.go
```

📜 *Sample Output*

```bash
1024
9.261000000000001
0.25
```

🏗 **N-Queens**

The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

📌 *Example Execution*

```bash
go run cmd/n_queens.go
```

📜 *Sample Output*

```bash
.Q..
...Q
Q...
..Q.

..Q.
Q...
...Q
.Q..
```

🏗 **Maximum Subarray**

Given an integer array nums, find the subarray with the largest sum, and return its sum.

📌 *Example Execution*

```bash
go run cmd/maximum_sub_array.go
```

📜 *Sample Output*

```bash
6
```

🏗 **Spiral Matrix**

Given an m x n matrix, return all elements of the matrix in spiral order.

📌 *Example Execution*

```bash
go run cmd/spiral_matrix.go
```

📜 *Sample Output*

```bash
[1 2 3 6 9 8 7 4 5]
```

🏗 **Edit Distance**

Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character

📌 *Example Execution*

```bash
go run cmd/edit_distance.go
```

📜 *Sample Output*

```bash
3
5
```

🏗 **Jump Game**

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

📌 *Example Execution*

```bash
go run cmd/jump_game.go
```

📜 *Sample Output*

```bash
true
false
```

🏗 **Merge Intervals**

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

📌 *Example Execution*

```bash
go run cmd/merge_intervals.go
```

📜 *Sample Output*

```bash
[[1 6] [8 10] [15 18]]
```

🏗 **Insert Interval**

You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Note that you don't need to modify intervals in-place. You can make a new array and return it.

📌 *Example Execution*

```bash
go run cmd/insert_interval.go
```

📜 *Sample Output*

```bash
[[1 5] [6 9]]
[[1 2] [3 10] [12 16]]
```

🏗 **Spiral Matrix II**

Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

📌 *Example Execution*

```bash
go run cmd/spiral_matrix2.go
```

📜 *Sample Output*

```bash
[1 2 3]
[8 9 4]
[7 6 5]
```

🏗 **Gray Code**

An n-bit gray code sequence is a sequence of 2n integers where:

Every integer is in the inclusive range [0, 2n - 1],
The first integer is 0,
An integer appears no more than once in the sequence,
The binary representation of every pair of adjacent integers differs by exactly one bit, and
The binary representation of the first and last integers differs by exactly one bit.
Given an integer n, return any valid n-bit gray code sequence.

📌 *Example Execution*

```bash
go run cmd/gray_code.go
```

📜 *Sample Output*

```bash
[0 1 3 2]
[0 1 3 2 6 7 5 4]
```

🏗 **Rotate List**

Given the head of a linked list, rotate the list to the right by k places.

📌 *Example Execution*

```bash
go run cmd/rotate_list.go 
```

📜 *Sample Output*

```bash
4 5 1 2 3
```

🏗 **Unique Paths**

There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.

📌 *Example Execution*

```bash
go run cmd/unique_paths.go
```

📜 *Sample Output*

```bash
28
```

🏗 **Unique Paths II**

You are given an m x n integer array grid. There is a robot initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

An obstacle and space are marked as 1 or 0 respectively in grid. A path that the robot takes cannot include any square that is an obstacle.

Return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The testcases are generated so that the answer will be less than or equal to 2 * 109.

📌 *Example Execution*

```bash
go run cmd/unique_paths2.go
```

📜 *Sample Output*

```bash
2
```

🏗 **Set Matrix Zeroes**

Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.

📌 *Example Execution*

```bash
go run cmd/set_matrix_zeros.go
```

📜 *Sample Output*

```bash
[[1 0 1] [0 0 0] [1 0 1]]
```

🏗 **Search a 2D Matrix**

You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

📌 *Example Execution*

```bash
go run cmd/search_a_2d_matrix.go
```

📜 *Sample Output*

```bash
true
false
```

🏗 **Sort Colors**

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

📌 *Example Execution*

```bash
go run cmd/sort_colors.go
```

📜 *Sample Output*

```bash
[0 0 1 1 2 2]
```

🏗 **Combinations**

Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.

📌 *Example Execution*

```bash
go run cmd/combinations.go
```

📜 *Sample Output*

```bash
[[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
```

🏗 **Subsets**

Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

📌 *Example Execution*

```bash
go run cmd/subsets.go
```

📜 *Sample Output*

```bash
[[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
```

🏗 **Word Search**

Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

📌 *Example Execution*

```bash
go run cmd/word_search.go
```

📜 *Sample Output*

```bash
true
```

🏗 **Remove Duplicates from Sorted Array II**

Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

📌 *Example Execution*

```bash
go run cmd/remove_duplicates_from_sorted_array2.go
```

📜 *Sample Output*

```bash
5
[1 1 2 2 3]
```

🏗 **Minimum Path Sum**

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

📌 *Example Execution*

```bash
go run cmd/minimum_path_sum.go
```

📜 *Sample Output*

```bash
7
```

🏗 **Partition List**

Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

📌 *Example Execution*

```bash
go run cmd/partition_list.go
```

📜 *Sample Output*

```bash
Original: 1 4 3 2 5 2 
Partitioned: 1 2 2 4 3 5 
```

🏗 **Decode Ways**

You have intercepted a secret message encoded as a string of numbers. The message is decoded via the following mapping:

"1" -> 'A'

"2" -> 'B'

📌 *Example Execution*

```bash
go run cmd/decode_ways.go
```

📜 *Sample Output*

```bash
Input: 226
Number of ways to decode: 3
```

🏗 **Reverse Linked List II**

Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

📌 *Example Execution*

```bash
go run cmd/reverse_linked_list2.go
```

📜 *Sample Output*

```bash
Original list:
1 2 3 4 5 
Reversed from position 2 to 4:
1 4 3 2 5 
```

🏗 **Search in Rotated Sorted Array II**

There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

You must decrease the overall operation steps as much as possible.

📌 *Example Execution*

```bash
go run cmd/search_in_rotated_sorted_array2.go
```

📜 *Sample Output*

```bash
Array: [2 5 6 0 0 1 2]
Target: 0
Target found in array.
```

🏗 **Unique Binary Search Trees II**

Given an integer n, return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n. Return the answer in any order.

📌 *Example Execution*

```bash
go run cmd/unique_binary_search_trees2.go
```

📜 *Sample Output*

```bash
Generating all unique BSTs for n = 3
Total trees generated: 5
Tree 1: 1,nil,2,nil,3,nil,nil
Tree 2: 1,nil,3,2,nil,nil,nil
Tree 3: 2,1,nil,nil,3,nil,nil
Tree 4: 3,1,nil,2,nil,nil,nil
Tree 5: 3,2,1,nil,nil,nil,nil
```

🏗 **Unique Binary Search Trees**

Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.

📌 *Example Execution*

```bash
go run cmd/unique_binary_search_trees.go
```

📜 *Sample Output*

```bash
Input: 3
Number of unique BSTs: 5
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
