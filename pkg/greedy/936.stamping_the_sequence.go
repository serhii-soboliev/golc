package greedy

/*
936. Stamping The Sequence
https://leetcode.com/problems/stamping-the-sequence/
*/

var QUESTION_MARK byte = '?'

func movesToStamp(stamp string, target string) []int {
	n := len(target)
	stampLen := len(stamp)
	cur := make([]byte, n) 
	for i:=0; i<n; i++  {
		cur[i] = target[i]
	}
	
	matchStamp := func(pos int) bool {
		ok := false

		for i:=0; i<stampLen; i++ {
			if cur[i + pos] != QUESTION_MARK {
				if cur[i + pos] != stamp[i] { 
					return false 
				}
				ok = true
			}		
		}

		return ok
	}

	replaceWithQuesitonMarks := func(pos int) {
		for i:=0; i<stampLen; i++ {
			cur[i + pos] = QUESTION_MARK	
		}
	}

	answer := []int{}
	isStamped := make([]bool, n) 

	for {
		found := false
		for i:=0; i<n-len(stamp)+1; i++ {
			if !isStamped[i] && matchStamp(i) {
				replaceWithQuesitonMarks(i)
				isStamped[i] = true
				found = true
				answer = append(answer, i)
			}
		}
		if !found { break }
	}

	for i:=0; i<n; i++ {
		if cur[i] != QUESTION_MARK {
			return []int{}
		}
	}

	for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}

	return answer
}

func MovesToStamp(stamp string, target string) []int {
	return movesToStamp(stamp, target)
}