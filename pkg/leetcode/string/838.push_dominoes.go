package string

/*
838. Push Dominoes
https://leetcode.com/problems/push-dominoes/
*/

func pushDominoes(dominoes string) string {
	dominoesInByteSlice := []byte(dominoes)
	lastIsRight := false
	lastIndex := 0
	for i := 0; i < len(dominoesInByteSlice); i++ {
		switch dominoesInByteSlice[i] {
		case 'L':
			if lastIsRight {
				count := (i - lastIndex + 1) / 2 - 1
				for m := lastIndex+1; m < lastIndex+1+count; m++ {
					dominoesInByteSlice[m] = 'R'
				}
				for m := i-1; m > i-1-count; m-- {
					dominoesInByteSlice[m] = 'L'
				}
			} else {
				for m := lastIndex; m < i; m++ {
					dominoesInByteSlice[m] = 'L'
				}
			}
			lastIsRight = false
			lastIndex = i+1
		case 'R':
			if lastIsRight {
				for m := lastIndex; m < i; m++ {
					dominoesInByteSlice[m] = 'R'
				}
			}
			lastIsRight = true
			lastIndex = i
		}
	}
	
	if lastIsRight {
		for m := lastIndex; m < len(dominoesInByteSlice); m++ {
			dominoesInByteSlice[m] = 'R'
		}
	}
	
	return string(dominoesInByteSlice)
}