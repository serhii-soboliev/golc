package backtracking

import "fmt"

/*
816. Ambiguous Coordinates
https://leetcode.com/problems/ambiguous-coordinates/
*/
func ambiguousCoordinates(st string) []string {
	res := []string{}

	isNotLegal := func(s string) bool {
		return !(len(s) == 1 || s[0] != '0' || s[len(s)-1] != '0')
	}

	makeCoordinates := func(s string) []string {
		switch {
			case len(s) == 1 || s[len(s)-1] == '0':
				return []string{s}
			case s[0] == '0':
				return []string{fmt.Sprintf("0.%s", s[1:])}
			default:
				result := make([]string, 1, len(s))
				result[0] = s
				for i := 1; i < len(s); i++ {
					result = append(result, fmt.Sprintf("%s.%s", s[:i], s[i:]))
				}
				return result
			}
	}

	for i:=2; i<len(st)-1; i++ {
		left, right := st[1:i], st[i:len(st)-1]
		if isNotLegal(left) || isNotLegal(right) {
			continue
		}
		leftC, rightC := makeCoordinates(left), makeCoordinates(right)
		for _, l := range leftC {
			for _, r := range rightC {
				res = append(res, []string{fmt.Sprintf("(%s, %s)", l, r)} ...)
			}
		}
	}

	return res
}

func AmbiguousCoordinates(s string) []string {
	return ambiguousCoordinates(s)
}