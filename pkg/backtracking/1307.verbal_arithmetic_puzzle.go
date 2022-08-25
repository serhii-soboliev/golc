package backtracking

/*
1307. Verbal Arithmetic Puzzle
https://leetcode.com/problems/verbal-arithmetic-puzzle/
*/

func removeIdx[T any](a []T, i int) []T{ 
	res := []T{}
	res = append(append(res, a[:i] ...), a[i+1:]...)
	return res
}

func Combinations[T any](a []T, k int) (result [][]T) {
	
	var combBacktrack func(pos int,l int, current []T)

	combBacktrack = func(pos int, l int, current []T) {
		if l == 0 {
			result = append(result, current)
			return
		}
		currendIdx := k - l
		for i:=pos; i<=len(a)-l; i++ {
		    cpy := make([]T, len(current))
		    copy(cpy, current)
			cpy[currendIdx] = a[i]
			combBacktrack(i+1, l-1, cpy)
		}
	}

	combBacktrack(0, k, make([]T, k))

	return
}


func GetAllBijections(chars []rune, digits []int) (result []map[rune]int) {

	var helper func(chars []rune, digits []int) (result []map[rune]int) 

	helper = func(chars []rune, digits []int) (result []map[rune]int) {
		if len(chars) == 0 {
			result = append(result, map[rune]int{})
			return 
		}
		char := chars[0]
		newChars := removeIdx(chars, 0)
		for i, digit := range digits {
			newDigits := removeIdx(digits, i)
			curResults := helper(newChars, newDigits)
			for _, curRes := range curResults {
				curRes[char] = digit
				result = append(result, curRes)
			}
		}

		return 
	}

	k := len(chars)
	combinatons := Combinations(digits, k)

	for _, combination := range combinatons {
		result = append(result, helper(chars, combination)...)
	}

	return
}

func isSolvable(words []string, result string) bool {

	getCharacters := func(allwords []string) (res []rune) {
		resMap := make(map[rune]bool)
		for _, word := range allwords {
			for _, b := range word {
				resMap[b] = true
			}	
		}
		for k := range resMap {
			res = append(res, k)
		}
		return 
	}
	
	calculateValue := func(word string, b map[rune]int) int {
		sum := 0
		for i:=len(word)-1; i>=0; i-- {
			d := rune(word[i])
			sum = sum * 10 + b[d]
		}
		return sum
	}

	chars := getCharacters(append(words, result))
	digits := []int{0,1,2,3,4,5,6,7,8,9}

	bijections := GetAllBijections(chars, digits)

	for _, b := range bijections {
		sum := 0
		for _, word := range words {
			sum += calculateValue(word, b)
		}
		if sum == calculateValue(result, b) {
			return true
		}
	}

	return false
}

func IsSolvable(words []string, result string) bool {
	return isSolvable(words, result)
}

