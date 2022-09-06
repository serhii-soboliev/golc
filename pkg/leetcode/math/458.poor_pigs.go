package math

import "math"

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

func poorPigsM(buckets int, minutesToDie int, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))	
}

func PoorPigsM(buckets int, minutesToDie int, minutesToTest int) int {
	return poorPigsM(buckets, minutesToDie, minutesToTest)
}

func PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	return poorPigs(buckets, minutesToDie, minutesToTest)
}