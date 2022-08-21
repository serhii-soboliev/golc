package greedy


/*
936. Stamping The Sequence
https://leetcode.com/problems/stamping-the-sequence/
*/

var QUESTION_MARK byte = '?'

func movesToStampCovers(stamp string, target string) []int {

	sLen, tLen := len(stamp), len(target)

	bTarget := []byte(target)
	bStamp := []byte(stamp)
	
	generateCovers := func() (result [][]byte) {

		for i:=0; i<sLen; i++ {
			for j:=0; j<sLen-i; j++ {
				cover := make([]byte, sLen)
				for k:=0; k<i; k++ {
					cover[k] = QUESTION_MARK
				}
				for k:=i; k<sLen-j; k++ {
					cover[k] = bStamp[k]
				}
				for k:=sLen-j; k<sLen; k++ {
					cover[k] = QUESTION_MARK
				}
				result = append(result, cover)
			}
		}

		return
	}

	replaceWithQuesitonMarks := func(pos int) {
		for i := 0; i < sLen; i++ {
			bTarget[i+pos] = QUESTION_MARK
		}
	}

	matchCover := func(pos int, cover []byte) bool {
		for i := 0; i < sLen; i++ {
			if bTarget[i+pos] != cover[i] { return false}
		}
		return true
	}

	isStamped := func () bool {
		for _, v := range bTarget {
			if v != QUESTION_MARK {
				return false
			}
		}
		return true	
	}

	covers := generateCovers()
	answer := []int{}

	p := tLen - sLen
	for !isStamped() {
		found := false
		for i:=p; i>=0; i-- {
			for _, cover := range covers {
				if matchCover(i, cover) {
					found = true
					answer = append([]int{i}, answer...)
					replaceWithQuesitonMarks(i)
				}
			}
		}	
		if !found  { return []int{} }
	}

	return answer
}

func movesToStamp(stamp string, target string) []int {

	min := func(a, b int) int {
		if a > b { return b} 
		return a
	}
	
	max := func(a, b int) int {
		if a < b { return b} 
		return a
	}

	type Node struct {
		todo map[int]bool
		made map[int]bool
	}

	m, n := len(stamp), len(target)
	answer := []int{}
	isStamped := make([]bool, n)
	q := []int{}
	a := []Node{}

	for i:=0; i<n-m+1; i++ {
		todo := make(map[int]bool)
		made := make(map[int]bool)
		for j:=0; j<m; j++ {
			v := i+j
			if target[v] == stamp[j]  {
				made[v] = true
			} else {
				todo[v] = true
			}
		}
		a = append(a, Node{todo, made})
		if len(todo) == 0 {
			answer = append([]int{i}, answer...)
			for j:=i; j<i+m; j++ {
				if isStamped[j] { continue }
				q = append(q, j)
				isStamped[j] = true	
			}
		}
	}

	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		for j := max(0, i-m+1); j<= min(n-m, i); j++ {
			node := a[j]
			if _, ok := node.todo[i]; ok {
				delete(node.todo, i)
				if len(node.todo) == 0 {
					answer = append([]int{j}, answer...)
					for k := range node.made {
						if isStamped[k] { continue } 
						q = append(q, k)
						isStamped[k] = true
					}
				}
			}
		}
	}

	for _, v := range isStamped {
		if !v { return []int{}}	
	}
	
	return answer
}

func movesToStampGreedy(stamp string, target string) []int {
	n := len(target)
	stampLen := len(stamp)
	cur := make([]byte, n)
	for i := 0; i < n; i++ {
		cur[i] = target[i]
	}

	matchStamp := func(pos int) bool {
		ok := false

		for i := 0; i < stampLen; i++ {
			if cur[i+pos] != QUESTION_MARK {
				if cur[i+pos] != stamp[i] {
					return false
				}
				ok = true
			}
		}

		return ok
	}

	replaceWithQuesitonMarks := func(pos int) {
		for i := 0; i < stampLen; i++ {
			cur[i+pos] = QUESTION_MARK
		}
	}

	answer := []int{}
	isStamped := make([]bool, n)

	for {
		found := false
		for i := 0; i < n-len(stamp)+1; i++ {
			if !isStamped[i] && matchStamp(i) {
				replaceWithQuesitonMarks(i)
				isStamped[i] = true
				found = true
				answer = append(answer, i)
			}
		}
		if !found {
			break
		}
	}

	for i := 0; i < n; i++ {
		if cur[i] != QUESTION_MARK {
			return []int{}
		}
	}

	// for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
	// 	answer[i], answer[j] = answer[j], answer[i]
	// }

	return answer
}

func MovesToStampGreedy(stamp string, target string) []int {
	return movesToStampGreedy(stamp, target)
}

func MovesToStamp(stamp string, target string) []int {
	return movesToStamp(stamp, target)
}

func MovesToStampCovers(stamp string, target string) []int {
	return movesToStampCovers(stamp, target)
}
