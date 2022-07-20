package backtracking

/*
784. Letter Case Permutation
https://leetcode.com/problems/letter-case-permutation/
*/
func isUpperCaseLetter(symbol byte) bool {
	return symbol >= 'A' && symbol <= 'Z'
}

func isDigit(symbol byte) bool {
	return '0' <= symbol && symbol <= '9'
}

func permute(word []byte, idx int, ans *[]string) {
	if idx == len(word) {
		*ans = append(*ans, string(word))
		return
	}
	currentSymbol := word[idx]
	newIdx := idx + 1
	if isDigit(currentSymbol) {
		permute(word, newIdx, ans)
		return
	}
	var newSymbol byte
	if isUpperCaseLetter(currentSymbol) {
		newSymbol = currentSymbol + 32
	} else {
		newSymbol = currentSymbol - 32
	}
	var oldWord = make([]byte, len(word))
	copy(oldWord, word)
	var newWord = append(word[:idx], append([]byte{newSymbol}, word[newIdx:]...)...)
	permute(oldWord, newIdx, ans)
	permute(newWord, newIdx, ans)
}

func letterCasePermutation(S string) []string {
	ans, word := []string{}, []byte(S)
	permute(word, 0, &ans)
	return ans
}

func LetterCasePermutation(s string) []string {
	return letterCasePermutation(s)
}
