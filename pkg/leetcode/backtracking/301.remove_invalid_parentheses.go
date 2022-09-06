package backtracking

/*
301. Remove Invalid Parentheses
https://leetcode.com/problems/remove-invalid-parentheses/
*/

func removeInvalidParentheses(s string) []string {
	result := NewSet[string]()
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
					result.Reset() 
					result.Add(currentExpression)
				}  else if removeCount == removeMinimum {
					result.Add(currentExpression)
				}
				
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

	return result.ToSlice()
}

type Set[T comparable] struct {
	m map[T]bool
}

func (set *Set[T]) Add(s T) {
	set.m[s] = true
}

func (set *Set[T]) Reset() {
	set.m = make(map[T]bool)
}

func (set *Set[T]) ToSlice() []T {
	result := []T{}
	for k := range set.m {
		result = append(result, k)	
	}
	return result	
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{ m: make(map[T]bool) }
}

func RemoveInvalidParentheses(s string) []string {
	return removeInvalidParentheses(s)
}
