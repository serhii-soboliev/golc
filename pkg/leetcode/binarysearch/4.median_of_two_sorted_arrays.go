package binarysearch

import "math"

/*
4. Median of Two Sorted Arrays
https://leetcode.com/problems/median-of-two-sorted-arrays/
*/

func findMedianSortedArrays(a []int, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	m, n := len(a), len(b)
	low, high, medianPos, total := 0, len(a), (m+n+1)/2, m+n

	for low <= high  {
		cut1 := (low+high)/2;
        cut2 := medianPos - cut1;
		l1 := ternary(cut1 == 0, math.MinInt, a, cut1-1);
        l2 := ternary(cut2 == 0, math.MinInt, b, cut2-1);
        r1 := ternary(cut1 == m, math.MaxInt, a, cut1);
        r2 := ternary(cut2 == n, math.MaxInt, b, cut2);
		if l1 <= r2 && l2 <= r1 {
			if total % 2 == 1 {
				return max(l1, l2)
			} else {
				return (max(l1,l2)+min(r1,r2))/2;
			}
		} else if l1 > r2 {
			 high = cut1-1
		} else { 
			low = cut1+1
		}
	}
	return 0
}

func ternary(exp bool, a int, nums []int, idx int) int {
	if exp {
		return a
	}
	return nums[idx]
}

func min(a, b int) float64 {
	if a < b {
		return float64(a)
	}
	return float64(b)
}

func max(a, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}