package backtracking

/*
301. Remove Invalid Parentheses
https://leetcode.com/problems/remove-invalid-parentheses/
*/
type Set struct {
	m map[string]bool
}

func (set *Set) Add(s string) {
	set.m[s] = true
}

func (set *Set) Reset() {
	set.m = make(map[string]bool)
}

func (set *Set) ToSlice() []string {
	result := []string{}
	for k := range set.m {
		result = append(result, k)	
	}
	return result	
}

func removeInvalidParentheses(s string) []string {
	resMap := Set { m: make(map[string]bool) }
	n := len(s)
	removeMinimum := n + 1 

	var backtracking func(idx int, leftCount int, rightCount int, removeCount int, currentExpression string) 
    
	backtracking = func(idx int, leftCount int, rightCount int, removeCount int, currentExpression string) { 
		if rightCount > leftCount {
			return
		}
		if idx == n {
			if leftCount == rightCount {
				if removeCount < removeMinimum {
					removeMinimum = removeCount
					resMap.Reset()
				} 
				resMap.Add(currentExpression)
			}
			return
		}
		curSymbol := string(rune(s[idx]))
		newCurrentExpression := currentExpression + curSymbol
		if curSymbol != "(" && curSymbol != ")" {
			backtracking(idx + 1, leftCount, rightCount, removeCount, newCurrentExpression)
		} else if curSymbol == "(" {
			backtracking(idx + 1, leftCount+1, rightCount, removeCount, newCurrentExpression)	
			backtracking(idx + 1, leftCount, rightCount, removeCount+1, currentExpression)	
		} else {
			backtracking(idx + 1, leftCount, rightCount+1, removeCount, newCurrentExpression)	
			backtracking(idx + 1, leftCount, rightCount, removeCount+1, currentExpression)		
		}
	}

	backtracking(0, 0, 0, 0, "")

	return resMap.ToSlice()
}

func RemoveInvalidParentheses(s string) []string {
	return removeInvalidParentheses(s)
}
