package hashmap

/*
2364. Count Number of Bad Pairs
https://leetcode.com/problems/count-number-of-bad-pairs/
*/

func countBadPairs(nums []int) int64 {
	k := len(nums)
	total := int64(k*(k-1) / 2)
	var good int64
	d := make(map[int64]int64)
	for i, v := range nums {
		diff := int64(i - v)
		if v, ok := d[diff]; ok {
			good += v
			d[diff] += 1
		} else {
			d[diff] = 1
		}	
	}

	return total - good
}