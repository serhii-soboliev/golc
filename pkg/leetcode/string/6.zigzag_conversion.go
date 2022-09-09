package string

/*
6. Zigzag Conversion
https://leetcode.com/problems/zigzag-conversion/
*/

func convert(s string, n int) string {
	if n == 1 {
		return s
	}
	l := len(s)
	result := []byte{}
	for k:=0; k<n; k++ {
		if k == 0 || k == n-1 {
			for i:=k; i<l; i+=(n-1)*2 {
				result = append(result,s[i])
			}
		} else {
			for i,j := k,k+(n-k-1)*2; ;i,j = i+(n-1)*2,j+(n-1)*2 {
				if i >= l {
					break
				}
				result = append(result,s[i])
				if j >= l {
					break
				}
				result = append(result,s[j])
			}
		}
	}
	return string(result)
}

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}
