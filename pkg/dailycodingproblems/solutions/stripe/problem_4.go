package stripe

/*
This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.

Leetcode 
41. First Missing Positive
https://leetcode.com/problems/first-missing-positive/

*/

func firstMissingPositive(nums []int) int {

	abs := func(a int) int {
		if a<0 {
			return -a
		}
		return a
	}
	
	l := len(nums)
	for i, v := range nums {
		if v < 0 {
			nums[i] = 0
		}
	}

	for _, v := range nums {
		av := abs(v)
		if av > l || v == 0 {
			continue
		}
		if nums[av-1] == 0 {
			nums[av-1] = -av
		} else {
			nums[av-1] = -abs(nums[av-1])	
		}
	}

	for i:=1; i<=l; i++ {
		if nums[i-1] >= 0 {
			return i
		}
	}
	return l+1
}

func FirstMissingPositive(nums []int) int {
	return firstMissingPositive(nums)
}