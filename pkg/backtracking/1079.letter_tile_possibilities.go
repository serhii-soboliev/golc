package backtracking


/*
1079. Letter Tile Possibilities
https://leetcode.com/problems/letter-tile-possibilities/
*/

func numTilePossibilities(tiles string) int {

	getLettersCounter := func() map[rune]int {
		res := make(map[rune]int)
		for _, v := range tiles {
			res[v] += 1
		}
		return res
	}

	letterCounter := getLettersCounter()
	result := 0
	n := len(tiles)

	var backtrack func(pos int)

	backtrack = func(pos int) {
		if pos > 0 {
			result += 1
			if pos == n {
				return 
			}
		}	
		for k, v := range letterCounter {
			if v <= 0 { continue }
			letterCounter[k] -= 1
			backtrack(pos + 1)
			letterCounter[k] += 1
		}
	}

	backtrack(0)
	return result
}

func NumTilePossibilities(tiles string) int {
	return numTilePossibilities(tiles)
}