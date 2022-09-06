package slidingwindow

/*
3. Longest Substring Without Repeating Characters
https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/

func lengthOfLongestSubstring(s string) int {
	charCounter := make(map[byte]int)
	curResult := 0
	result := 0
	substringStart := 0

	for i := 0; i < len(s); i++ {
		curChar := s[i]

		for charCounter[curChar] > 0 {
			charCounter[s[substringStart]] -= 1
			substringStart += 1
			curResult -= 1
		}

		charCounter[curChar] = 1
		curResult += 1
		if curResult > result {
			result = curResult
		}
	}

	return result
}
