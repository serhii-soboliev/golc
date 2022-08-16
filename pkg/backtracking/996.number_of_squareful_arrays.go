package backtracking

import "fmt"

/*
996. Number of Squareful Arrays
https://leetcode.com/problems/number-of-squareful-arrays/
*/

func numSquarefulPerms(nums []int) int {
	n := len(nums)
	used := make([]byte, n)

	squarefulPerms := make(map[string]struct{})

	toString := func(a []int) string {
		res := ""
		for _, v := range a {
			res = fmt.Sprintf("%s%d",res,v)
		}
		return res
	}

	isPerfectSquare := func(a, b int) bool {
		x := a + b
		lo, hi := 1, x
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
		l := len(a)
		if len(a) == n { 
			squarefulPerms[toString(a)] = struct{}{}
			return
		} 
		for i, v := range(nums) {
			if used[i] == 1 { continue }
			if l > 0 && !isPerfectSquare(a[l-1], v) { continue }
			u[i] = 1
			a = append(a, v)
			backtrack(a, u)
			u[i] = 0
			a = a[:l]
		}		
	}

	backtrack([] int {}, used)
	return len(squarefulPerms)
}

func NumSquarefulPerms(nums []int) int {
	return numSquarefulPerms(nums)
}