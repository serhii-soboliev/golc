package dynamicprogramming

/*
718. Maximum Length of Repeated Subarray
https://leetcode.com/problems/maximum-length-of-repeated-subarray/
*/

func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	memo := make([][]int, n+1)
	for i:=0; i<n+1; i++ {
		memo[i] = make([]int, m+1)
	}
	result := 0
	for i:=n-1; i>=0; i-- {
		for j:=m-1; j>=0; j-- {
			if nums1[i] == nums2[j] {
				memo[i][j] = memo[i+1][j+1] + 1
				if result < memo[i][j] {
					result = memo[i][j]
				}
			}
		}
	}
	return result

}