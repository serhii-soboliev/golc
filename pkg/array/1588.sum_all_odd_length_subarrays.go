package array

/*
1588. Sum of All Odd Length Subarrays
https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
*/

func sumOddLengthSubarrays(arr []int) int {

	n := len(arr)
   
	buildPrefixSum := func() []int {
		result := make([]int, n+1) 
		for i:=0; i<n; i++ {
			result[i+1] = arr[i] + result[i]
		}
		return result
	}

	prefixSum := buildPrefixSum() 
	sum := 0

	for i:=0; i<n; i++ {
		for j:=1; i+j<=n; j+=2 {
			curSum := prefixSum[i+j] - prefixSum[i]
			sum += curSum
		}
	}
	return sum

}