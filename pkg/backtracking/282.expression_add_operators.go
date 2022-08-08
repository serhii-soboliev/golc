package backtracking

import (
	"fmt"
)

/*
282. Expression Add Operators
https://leetcode.com/problems/expression-add-operators/
*/

func addOperators(num string, target int) []string {
	result := [] string {}
	n := len(num)

	var backtrack func(idx int, prevO int,currO int, value int, ops []string)
	
	backtrack = func(idx int, prevO, currO, value int, ops []string) {
		if idx == n {
			if value == int(target) && currO == 0 {
				subresult := ""
				for j:=1;j<len(ops);j++ {
					subresult += ops[j]
				}
				result = append(result, subresult)
			}
			return
		}
		currO = currO*10 + int(num[idx] - '0')
		strCurrO := fmt.Sprint(currO)

		if currO > 0 {
			backtrack(idx + 1, prevO, currO, value, ops)
		}

		opsLen := len(ops)
		ops = append(ops, "+", strCurrO)
		backtrack(idx + 1, currO, 0, value + currO, ops)
		ops = ops[:opsLen]

		if len(ops) > 0 {
			ops = append(ops,"-", strCurrO)
			backtrack(idx + 1, -currO, 0, value - currO, ops)
			ops = ops[:opsLen]

			ops = append(ops, "*", strCurrO)
			backtrack(idx + 1, prevO * currO, 0, value - prevO + (currO * prevO), ops)
			ops = ops[:opsLen]
		}
	}	

	backtrack(0, 0, 0, 0, []string{})
	return result
}

func AddOperators(num string, target int) []string {
	return addOperators(num, target)
}