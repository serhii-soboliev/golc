package backtracking

import (
	"math"
)

/*
691. Stickers to Spell Word
https://leetcode.com/problems/stickers-to-spell-word/
*/

var LETTERS_COUNT = 26
var START_LETTER = 'a'

func buildLetterHistogram(s string) []rune {
	result := make([]rune, LETTERS_COUNT)
	for _, r := range s {
		result[r-START_LETTER] += 1
	}
	return result
}

func isMoreValuable(k int, n int, historgram [][]rune, target string) bool {
	for _, r := range target {
		i := r - START_LETTER
		if historgram[n][i] > historgram[k][i] {
			return false
		}
	}
	return true
}

func buildStickersLetterHistogram(stickers []string, target string) [][]rune {
	stickersCount := len(stickers)
	primaryRuneHistogram := make([][]rune, stickersCount)
	for i, sticker := range stickers {
		primaryRuneHistogram[i] = buildLetterHistogram(sticker)
	}
	notValuableStickers := make(map[int]bool)
	for i := 0; i < stickersCount; i++ {
		for j := 0; j < stickersCount; j++ {
			if _, ok := notValuableStickers[i]; ok || i == j {
				continue
			}
			if isMoreValuable(i, j, primaryRuneHistogram, target) {
				notValuableStickers[j] = true
			}
		}
	}

	stickersRuneHistogram := [][]rune{}
	for i := 0; i < stickersCount; i++ {
		if _, ok := notValuableStickers[i]; !ok {
			stickersRuneHistogram = append(stickersRuneHistogram, primaryRuneHistogram[i])
		}
	}
	return stickersRuneHistogram
}

func copyHistogram(c []rune) []rune {
	cpy := make([]rune, len(c))
	copy(cpy, c)
	return cpy
}

func minStickersDP(stickers []string, target string) int {

	stickersLetterHistogram := buildStickersLetterHistogram(stickers, target)

	var dfs func(string, []rune) int

	dp := make(map[string]int)

	dfs = func(t string, sticker []rune) int {
		if v, ok := dp[t]; ok {
			return v
		}
		res := 1
		if len(sticker) == 0 { res = 0}
		remainingWord := ""
		for _, r := range target {
			if sticker[r-START_LETTER] == 0 { 
				remainingWord += string(r) 
				sticker[r-START_LETTER] -= 1
			}
		}
		if len(remainingWord) > 0 {
			used := math.MaxInt
			for _, h := range stickersLetterHistogram {
				if h[remainingWord[0] - byte(START_LETTER)] == 0 { continue }
				tempRes := dfs(remainingWord, copyHistogram(h))
				if tempRes < used { used = tempRes}
			}
			dp[remainingWord] = used
			res += used
		}

		return res
	}

	result := dfs(target, make([]rune, LETTERS_COUNT))
	if result == math.MaxInt { return -1 }
	return result
}

func minStickersBacktrack(stickers []string, target string) int {
	minUsedStickersNumber := math.MaxInt32
	stickersLetterHistogram := buildStickersLetterHistogram(stickers, target)
	stickersCount := len(stickersLetterHistogram)

	var backtracking func(int, []rune, int)

	backtracking = func(idx int, availableHistogram []rune, usedStickersNumber int) {
		if usedStickersNumber > minUsedStickersNumber {
			return
		}
		if idx == len(target) {
			if minUsedStickersNumber > usedStickersNumber {
				minUsedStickersNumber = usedStickersNumber
			}
			return
		}
		rIdx := rune(target[idx]) - START_LETTER
		if availableHistogram[rIdx] > 0 {
			availableHistogram[rIdx] -= 1
			backtracking(idx+1, availableHistogram, usedStickersNumber)
			availableHistogram[rIdx] += 1
		} else {
			for i := 0; i < stickersCount; i++ {
				if stickersLetterHistogram[i][rIdx] == 0 {
					continue
				}
				for j := 0; j < LETTERS_COUNT; j++ {
					availableHistogram[j] += stickersLetterHistogram[i][j]
				}
				backtracking(idx, availableHistogram, usedStickersNumber+1)
				for j := 0; j < LETTERS_COUNT; j++ {
					availableHistogram[j] -= stickersLetterHistogram[i][j]
				}
			}
		}
	}

	backtracking(0, make([]rune, LETTERS_COUNT), 0)

	if minUsedStickersNumber == math.MaxInt32 {
		return -1
	}
	return minUsedStickersNumber
}

func MinStickersDP(stickers []string, target string) int {
	return minStickersBacktrack(stickers, target)
}

func MinStickersBacktrack(stickers []string, target string) int {
	return minStickersBacktrack(stickers, target)
}
