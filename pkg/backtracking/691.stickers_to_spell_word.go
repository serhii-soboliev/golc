package backtracking

import (
	"math"
	"strings"
)

/*
691. Stickers to Spell Word
https://leetcode.com/problems/stickers-to-spell-word/
*/

func minStickers(stickers []string, target string) int {
	targetSymbolsMap := make(map[rune]int)
	for _, r := range target { targetSymbolsMap[r] += 1 }
	minUsedStickersNumber := math.MaxInt32

	var backtracking func(int, map[rune]int, int) 

	backtracking = func(idx int, availableSymbolsMap map[rune]int, usedStickersNumber int) {
		if idx == len(target) {
			if minUsedStickersNumber > usedStickersNumber { minUsedStickersNumber = usedStickersNumber}
			return
		} 
		r := rune(target[idx])
		if targetSymbolsMap[r] <= availableSymbolsMap[r] {
			backtracking(idx + 1, availableSymbolsMap, usedStickersNumber)
		} else if usedStickersNumber + 1 < minUsedStickersNumber {
			for _, sticker := range stickers {
				if !strings.Contains(sticker, string(r)) { continue }
				for _, s := range sticker {
					availableSymbolsMap[s] += 1
				}
				backtracking(idx, availableSymbolsMap, usedStickersNumber + 1)
				for _, s := range sticker {
					availableSymbolsMap[s] -= 1
				}
			}
		}
	}
	
	backtracking(0, make(map[rune]int), 0)

	if minUsedStickersNumber == math.MaxInt32 {	return -1} 
	return minUsedStickersNumber
}

func MinStickers(stickers []string, target string) int {
	return minStickers(stickers, target)
}