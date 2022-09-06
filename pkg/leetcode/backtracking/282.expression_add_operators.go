package backtracking

import (
	"fmt"
	"strconv"
)

/*
282. Expression Add Operators
https://leetcode.com/problems/expression-add-operators/
*/

func addOperators(num string, target int) []string {
    result := []string{}
    
    var backtrack func(string, string, int, int)
    
    backtrack = func(suffix, path string, val, last int) {
        if len(suffix) == 0 {
            if val == target {
                result = append(result, path)
            }
            return
        }
        
        for i := 1; i <= len(suffix); i++ {
            str := suffix[:i]
            digit, _ := strconv.Atoi(str)
            
            if str[0] == '0'&& len(str) > 1 {
                continue
            }
            
			// for the first call
            if path == "" {
                backtrack(suffix[i:], str, digit, digit)
            } else {
                backtrack(suffix[i:], path + "+" + str, val + digit, digit)
                backtrack(suffix[i:], path + "-" + str, val - digit, -digit)
                backtrack(suffix[i:], path + "*" + str, val - last + last * digit, last * digit)
            }
        }
    }
    
    backtrack(num, "", 0, 0)
    
    return result
}

func addOperators1(num string, target int) []string {
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

func AddOperators1(num string, target int) []string {
	return addOperators1(num, target)
}