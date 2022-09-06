package backtracking

/*
1255. Maximum Score Words Formed by Letters
https://leetcode.com/problems/maximum-score-words-formed-by-letters/
*/

func maxScoreWords(words []string, letters []byte, score []int) int {

	getLettersCounter := func() map[byte]int {
		res := make(map[byte]int)
		for _,v := range letters {
			res[v] += 1
		}
		return res
	}

	lettersCount := getLettersCounter() 

	getWordsLettersCountersAndScores := func() ([]map[byte]int, []int) {
		counters, scores := []map[byte]int{}, []int{}
		for _, word := range words {
			wordLetterCounter := make(map[byte]int)
			wordScore := 0
			bWord := []byte(word)
			isImpossibleWord := false
			for _, v := range bWord {
				wordScore += score[v - 'a']
				wordLetterCounter[v] += 1
				if wordLetterCounter[v] > lettersCount[v] {
					isImpossibleWord = true
				}
			}

			if wordScore > 0 || isImpossibleWord {
				counters = append(counters, wordLetterCounter)
				scores = append(scores, wordScore)
			}
		}
		return counters, scores
	}

	wordLettersCounters, wordScores := getWordsLettersCountersAndScores()

	canUse := func(wordCounter map[byte]int) bool {
		for k, v := range wordCounter {
			if v > lettersCount[k] {
				return false
			}
		}
		return true
	}

	use := func(wordCounter map[byte]int) {
		for k, v := range wordCounter {
			lettersCount[k] -= v
		}
	}

	revertUse := func(wordCounter map[byte]int) {
		for k, v := range wordCounter {
			lettersCount[k] += v
		}
	}

	valuableWordsCount := len(wordScores)
	maxScore := 0
	
	var backtrack func(pos int, currentScore int)
	
	backtrack = func(pos int, currentScore int) {

		if pos == valuableWordsCount {
			if currentScore > maxScore {
				maxScore = currentScore
			}
			return
		}

		currentWordLetterCounter := wordLettersCounters[pos]
		if canUse(currentWordLetterCounter) {
			use(currentWordLetterCounter)
			backtrack(pos+1, currentScore + wordScores[pos])
			revertUse(currentWordLetterCounter)
		}

		backtrack(pos + 1, currentScore)
		
	}

	backtrack(0, 0)

	return maxScore
}

func MaxScoreWords(words []string, letters []byte, score []int) int {
	return maxScoreWords(words, letters, score)
}