package uber

/*
// Problem 2
//
// This problem was asked by Uber.
//
// Given an array of integers, return a new array such that each element at index i of the
// new array is the product of all the numbers in the original array except the one at i.
//
// For example, if our input was [1, 2, 3, 4, 5],
// the expected output would be [120, 60, 40, 30, 24].
// If our input was [3, 2, 1], the expected output would be [2, 3, 6].
//
// Follow-up: what if you can't use division?
//
// https://leetcode.com/problems/product-of-array-except-self/
//
// O(N) Time Complexity
// O(N) Space Complexity, O(1) if not including output array
// N is the number of elements in the array
//
// With division, multiply every number to have a final product and then for each number divide from the final product

On leetcode: 
238. Product of Array Except Self
https://leetcode.com/problems/product-of-array-except-self/submissions/

*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	productArr := make([]int, n)
	productArr[0] = 1
	for i:=1; i<n; i++ {
		productArr[i] = productArr[i-1]*nums[i-1]
	}
	product := 1
	for i:=n-1; i>=0; i-- {
		productArr[i] = productArr[i] * product
		product *= nums[i]
	}

	return productArr
}