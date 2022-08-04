package math

/*
858. Mirror Reflection
https://leetcode.com/problems/mirror-reflection/
*/
func mirrorReflectionMath(p int, q int) int {
	for p % 2 == 0 && q % 2 == 0 {
		p /= 2
		q /= 2
	}
	return 1 - p%2 + q % 2
}

func MirrorReflectionMath(p int, q int) int {
	return mirrorReflectionMath(p,q)
}