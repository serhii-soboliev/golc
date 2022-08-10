package backtracking

/*
301. Remove Invalid Parentheses
https://leetcode.com/problems/remove-invalid-parentheses/
*/

func removeInvalidParentheses(s string) []string {
	resMap := make(map[string]bool)
	n := len(s)
	removeMinimum := n + 1 

	isOpenPar := func(s string) bool {
		return s == "("
	}

	isClosePar := func(s string) bool {
		return s == ")"
	}

	var backtracking func(idx int, leftCount int, rightCount int, removeCount int, currentExpression string) 
    
	backtracking = func(idx int, leftCount int, rightCount int, removeCount int, currentExpression string) { 
		if rightCount > leftCount {
			return
		}
		if idx == n {
			if leftCount == rightCount {
				if removeCount < removeMinimum {
					removeMinimum = removeCount
					resMap = map[string]bool{currentExpression:true}
				} else if removeCount == removeMinimum {
					resMap[currentExpression] = true
				}
			}
			return
		}
		curSymbol := string(rune(s[idx]))
		newCurrentExpression := currentExpression + curSymbol
		if !isOpenPar(curSymbol) && !isClosePar(curSymbol) {
			backtracking(idx + 1, leftCount, rightCount, removeCount, newCurrentExpression)
		} else if isOpenPar(curSymbol) {
			backtracking(idx + 1, leftCount+1, rightCount, removeCount, newCurrentExpression)	
			backtracking(idx + 1, leftCount, rightCount, removeCount+1, currentExpression)	
		} else {
			backtracking(idx + 1, leftCount, rightCount+1, removeCount, newCurrentExpression)	
			backtracking(idx + 1, leftCount, rightCount, removeCount+1, currentExpression)		
		}
	}

	backtracking(0, 0, 0, 0, "")
	result := []string{}
	for k := range resMap {
		result = append(result, k)	
	}
	return result
}

func RemoveInvalidParentheses(s string) []string {
	return removeInvalidParentheses(s)
}
