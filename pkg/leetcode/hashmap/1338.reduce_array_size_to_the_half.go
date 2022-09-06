package hashmap

import "sort"

/*
1338. Reduce Array Size to The Half
https://www.youtube.com/watch?v=UvAdMevcZbA
*/

func minSetSize(arr []int) int {
	m := make(map[int]int)
	
	for _, a := range arr {
		m[a] += 1
	}

	s := []int {}
	for _, v := range m {
		s = append(s, v)
	}

	sort.Slice(s, func(a, b int) bool {return s[b] < s[a]})

	sum := 0
	for i:=0; i<len(s); i++ {
		sum += s[i] 
		if sum >= len(arr) / 2 {
			return i+1
		}
	}
	return -1

}