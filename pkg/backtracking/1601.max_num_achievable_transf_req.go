package backtracking

/*
1601. Maximum Number of Achievable Transfer Requests
https://leetcode.com/problems/maximum-number-of-achievable-transfer-requests/
*/

func maximumRequests(n int, requests [][]int) int {
	buildingBalance := make([]int, n)
	reqCnt := len(requests)

	isBalanced := func() bool {
		for _,v := range buildingBalance {
			if v != 0 {
				return false
			}
		}
		return true
	}

	result := 0

	var backtrack func(pos int, fReqCount int)

	backtrack = func(pos, fReqCount int) {
		if pos > reqCnt {
			return
		}
		if isBalanced() && fReqCount > result{
			result = fReqCount
		}
		for i:=pos; i<reqCnt; i++ {
			req := requests[i]
			buildingBalance[req[0]] -= 1
			buildingBalance[req[1]] += 1
			backtrack(i+1, fReqCount+1)
			buildingBalance[req[0]] += 1
			buildingBalance[req[1]] -= 1
		}
	}

	backtrack(0,0)
	return result
}

func MaximumRequests(n int, requests [][]int) int {
	return maximumRequests(n, requests)
}
