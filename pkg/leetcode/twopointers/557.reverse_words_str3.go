package twopointers

/*
557. Reverse Words in a String III
https://leetcode.com/problems/reverse-words-in-a-string-iii/
*/

func reverseWords(s string) string {
	r := []byte(s)

	reverse := func(i, j int) {
		for i<j {
			r[i],r[j] = r[j], r[i]
			i+=1
			j-=1
		}
	}
	l := len(s)
	start, end  := 0, 0

	for end < l-1 {
		for r[end] != ' ' && end < l-1 {
			end+=1
		}
		if r[end] == ' ' {
			reverse(start, end-1)
		} else {
			reverse(start, end)	
		}
		end += 1
		start = end
		
	}
	return string(r)
}

func ReverseWords(s string) string {
	return reverseWords(s)
}