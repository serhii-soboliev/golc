package backtracking

import "math"

/*
679. 24 Game
https://leetcode.com/problems/24-game/
*/

var epsilon = math.Pow10(-6)

func judgePoint24(cards []int) bool {

	generateAllPossibleResults := func(a, b float64) []float64 {
		res := [] float64 {a+b, a*b, a-b, b-a}
	    if b != 0 {res = append(res, a/b) }
	    if a != 0 {res = append(res, b/a) }
		return res
	}

	isValid := func(a []float64) bool {
		return math.Abs(a[0] - 24) <= epsilon
	}

	cardsWithoutIndexes := func(a []float64, index1, index2 int) []float64 {
		r := []float64{}
		for i, v := range a {
			if i != index1 && i != index2 {
				r = append(r, v)
			}
		}
		return r
	}

	var backtrack func(a []float64) bool 

	backtrack = func(a []float64) bool {
		l := len(a)
		if l == 1 { return isValid(a) }
		for i:=0; i<l;i++ {
			for j:=i+1; j<l; j++ {
				possibleResults := generateAllPossibleResults(a[i], a[j])
				for _,v  := range possibleResults {
					newA := cardsWithoutIndexes(a, i, j)
					newA = append(newA, v)
					if backtrack(newA) { return true }
				}
			}
		}
		return false
	}
	
	fCards := []float64{}
	for _,v := range cards {
		fCards = append(fCards, float64(v))	
	}
	return backtrack(fCards)
}

func JudgePoint24(cards []int) bool {
	return judgePoint24(cards)
}