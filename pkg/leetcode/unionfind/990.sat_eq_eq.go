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

	for i := range u {
		u[i] = byte(i) 
	}

	for _, eq := range equations {
		if eq[1] == '=' {
			u[find(eq[0] - 'a')] = find(eq[3] - 'a'); 	
		}
	}

	for _, eq := range equations {
		if eq[1] == '!' && find(eq[0] - 'a') == find(eq[3] - 'a') {
			return false
		}
	}

	return true
}

func EquationsPossible(equations []string) bool {
	return equationsPossible(equations)
}