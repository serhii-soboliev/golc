package math


/*
458. Poor Pigs
https://leetcode.com/problems/poor-pigs/
*/
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	pigs := 0
	testNum := minutesToTest / minutesToDie + 1
	n := 1
	for n < buckets {
		pigs += 1
		n *= testNum
	}
	return pigs
}

func PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	return poorPigs(buckets, minutesToDie, minutesToTest)
}