package unionfind

/*
990. Satisfiability of Equality Equations
https://leetcode.com/problems/satisfiability-of-equality-equations/
*/


func equationsPossible(equations []string) bool {
	u := [26]byte{}

	var find func(x byte) byte

	find = func(x byte) byte{
		if u[x] == x {
			return x
		} else {
			return find(u[x])
		}
	}

	unite := func(x, y byte) {
		u[find(x - 'a')] = find(y - 'a')
	}

	isUnited := func(x, y byte) bool{
		return find(x - 'a') == find(y - 'a')
	}

	for i := range u {
		u[i] = byte(i) 
	}

	for _, eq := range equations {
		if eq[1] == '=' {
			unite(eq[0], eq[3]) 	
		}
	}

	for _, eq := range equations {
		if eq[1] == '!' && isUnited(eq[0], eq[3]) {
			return false
		}
	}

	return true
}

func EquationsPossible(equations []string) bool {
	return equationsPossible(equations)
}