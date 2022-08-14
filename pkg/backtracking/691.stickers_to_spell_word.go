package backtracking

import (
	"math"
)

/*
691. Stickers to Spell Word
https://leetcode.com/problems/stickers-to-spell-word/
*/

func runeHistogram(s string) []rune {
	result := make([]rune, 26)
	for _, r := range s {
		result[r - 'a'] += 1
	}
	return result 
}

func moreValuable(k int, n int, historgram [][]rune, target string) bool {
	for _, r := range target {
		i := r - 'a'
		if historgram[n][i] > historgram[k][i] { return false }
	}
	return true
}

func buildStickersLetterHistogram(stickers []string, target string) [][]rune {
	stickersCount := len(stickers)
	primaryRuneHistogram := make([][]rune, stickersCount)
	for i, sticker := range stickers {
		primaryRuneHistogram[i] = runeHistogram(sticker)
	}
	notValuableStickers := make(map[int]bool)
	for i:=0; i<stickersCount; i++ {
		for j:=0; j<stickersCount; j++ {
			if _, ok := notValuableStickers[i]; ok || i == j { continue }
			if moreValuable(i, j, primaryRuneHistogram, target) { 
				notValuableStickers[j] = true
			}	
		}
	}

	stickersRuneHistogram := [][]rune{}
	for i:=0; i<stickersCount; i++ {
		if _, ok := notValuableStickers[i]; !ok {
			stickersRuneHistogram = append(stickersRuneHistogram, primaryRuneHistogram[i])
		}
	}
	return stickersRuneHistogram
}

func minStickers(stickers []string, target string) int {
	minUsedStickersNumber := math.MaxInt32
	stickersRuneHistogram := buildStickersLetterHistogram(stickers, target)
	stickersCount := len(stickersRuneHistogram)

	var backtracking func(int, []rune, int) 

	backtracking = func(idx int, availableHistogram []rune, usedStickersNumber int) {
		if usedStickersNumber > minUsedStickersNumber { return }
		if idx == len(target) {
			if minUsedStickersNumber > usedStickersNumber { minUsedStickersNumber = usedStickersNumber}
			return
		} 
		rIdx := rune(target[idx]) - 'a'
		if availableHistogram[rIdx] > 0  {
			availableHistogram[rIdx] -= 1
			backtracking(idx + 1, availableHistogram, usedStickersNumber)
			availableHistogram[rIdx] += 1
		} else {
			for i:=0; i<stickersCount; i++ {
				if stickersRuneHistogram[i][rIdx] == 0 { continue }
				for j:=0; j<26; j++ {
					availableHistogram[j] += stickersRuneHistogram[i][j]
				}
				backtracking(idx, availableHistogram, usedStickersNumber + 1)
				for j:=0; j<26; j++ {
					availableHistogram[j] -= stickersRuneHistogram[i][j]
				}
			}
		}
	}
	
	backtracking(0, make([]rune, 26), 0)

	if minUsedStickersNumber == math.MaxInt32 {	return -1} 
	return minUsedStickersNumber
}

func MinStickers(stickers []string, target string) int {
	return minStickers(stickers, target)
}