package math

/*
858. Mirror Reflection
https://leetcode.com/problems/mirror-reflection/
*/

func gcd(a, b int) int {
	if a == 0 || a == b { return b }
	if b == 0 { return a }
	if a > b {
		return gcd(a - b, b)
	} else {
		return gcd(a, b - a)
	}
}

func lcm(a, b int) int {
	return a*b / gcd(a, b)
}

func mirrorReflectionLCM(p int, q int) int {
	lcm := lcm(p, q)
	m := lcm / p
	n := lcm / q
	return 1 + m % 2 - n % 2
}

func mirrorReflectionMath(p int, q int) int {
	for p % 2 == 0 && q % 2 == 0 {
		p /= 2
		q /= 2
	}
	return 1 - p % 2 + q % 2
}

func MirrorReflectionLCM(p int, q int) int {
	return mirrorReflectionLCM(p,q)
}

func MirrorReflectionMath(p int, q int) int {
	return mirrorReflectionMath(p,q)
}