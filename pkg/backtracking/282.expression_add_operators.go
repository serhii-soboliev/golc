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

	var backtrack func(idx int, prevOperand int64,
							 currentOperand int64, 
							 value int64, ops []string)
	
	backtrack = func(idx int, prevOperand, currentOperand, value int64, ops []string) {
		if idx == n {
			if value == int64(target) && currentOperand == 0 {
				subresult := ""
				for j:=1;j<len(ops);j++ {
					subresult += ops[j]
				}
				result = append(result, subresult)
			}
			return
		}
		idxChar := rune(num[idx]) - '0'
		currentOperand = currentOperand*10 + int64(idxChar)
		strCurentOperand := fmt.Sprint(currentOperand)

		if currentOperand > 0 {
			backtrack(idx + 1, prevOperand, currentOperand, value, ops)
		}

		ops = append(ops, "+")
		ops = append(ops, strCurentOperand)
		backtrack(idx + 1, currentOperand, 0, value + currentOperand, ops)
		ops = ops[:len(ops) - 2]

		if len(ops) > 0 {
			ops = append(ops, "-")
			ops = append(ops, strCurentOperand)
			backtrack(idx + 1, -currentOperand, 0, 
				value - currentOperand, ops)
			ops = ops[:len(ops) - 2]

			ops = append(ops, "*")
			ops = append(ops, strCurentOperand)
			backtrack(idx + 1, prevOperand * currentOperand, 0,
				value - prevOperand + (currentOperand * prevOperand), ops)
			ops = ops[:len(ops) - 2]
		}
	}	

	backtrack(0, 0, 0, 0, []string{})
	return result
}

func AddOperators(num string, target int) []string {
	return addOperators(num, target)
}