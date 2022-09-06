package backtracking

import (
	"fmt"
	"sort"
)

/*
996. Number of Squareful Arrays
https://leetcode.com/problems/number-of-squareful-arrays/
*/

func numSquarefulPerms(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	squarefulPermsCount := 0
	usedPerms := make(map[string]struct{})

	toString := func(a []int) string {
		res := ""
		for _, v := range a {
			res = fmt.Sprintf("%s%d",res,v)
		}
		return res
	}

	isUsed := func(a []int) bool {
		p := toString(a)
		_, ok := usedPerms[p]
		return ok	
	}

	markAsUsed := func(a []int) {
		p := toString(a)
		usedPerms[p] = struct{}{}	
	}

	isPerfectSquare := func(a, b int) bool {
		x := a + b
		lo, hi := 0, x
     	for lo <= hi {
        	mid := (lo + hi) / 2;
			if (mid * mid == x) {
				return true;
			} else if (mid * mid < x) {
				lo = mid + 1;
			} else {
				hi = mid - 1;
			}
    	}
    	return false;
	}

	var backtrack func ([]int, []byte)

	backtrack = func(a []int, u []byte) {
		if isUsed(a) { return }
		markAsUsed(a)
		l := len(a)
		if len(a) == n { 
			squarefulPermsCount += 1
			return
		} 
		for i, v := range(nums) {
			if u[i] == 1 { continue }
			if l > 0 && !isPerfectSquare(a[l-1], v) { continue }
			u[i] = 1
			a = append(a, v)
			backtrack(a, u)
			u[i] = 0
			a = a[:l]
		}		
	}

	backtrack([] int {}, make([]byte, n))
	return squarefulPermsCount
}

func NumSquarefulPerms(nums []int) int {
	return numSquarefulPerms(nums)
}