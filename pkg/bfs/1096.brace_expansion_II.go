package bfs

import (
	"sort"
	"strings"
	"unicode"
)

/*
1096. Brace Expansion II
https://leetcode.com/problems/brace-expansion-ii/
*/
var OPENING_BRACE = "{"
var CLOSING_BRACE = "}"
var COMMA = ","

func braceExpansionRecII(e string) []string {
	n := len(e)
	if n == 0 {
		return []string{""}
	}

	findFirstComma := func() (i int) {
		open := 0
		for i = 0; i < n && !(string(e[i]) == COMMA && open == 0); i++ {
			if string(e[i]) == OPENING_BRACE {
				open += 1
			}
			if string(e[i]) == CLOSING_BRACE {
				open -= 1
			}
		}
		return i
	}

	concatenate := func(a, b []string) (res []string) {
		resMap := make(map[string]bool)
		for _, v1 := range a {
			for _, v2 := range b {
				resMap[v1+v2] = true
			}
		}
		for k := range resMap {
			res = append(res, k)
		}
		sort.Strings(res)
		return
	}

	union := func(a, b []string) (res []string) {
		resMap := make(map[string]bool)
		for _, v := range a {
			resMap[v] = true
		}
		for _, v := range b {
			resMap[v] = true
		}
		for k := range resMap {
			res = append(res, k)
		}
		sort.Strings(res)
		return
	}

	firstComma := findFirstComma()

	if firstComma < n {
		first := braceExpansionRecII(e[0:firstComma])
		second := braceExpansionRecII(e[firstComma+1:])
		return union(first, second)
	}

	if string(e[0]) == OPENING_BRACE {
		open := 0
		var i int
		for i = 0; i < n && !(string(e[i]) == CLOSING_BRACE && open == 1); i++ {
			if string(e[i]) == OPENING_BRACE {
				open += 1
			}
			if string(e[i]) == CLOSING_BRACE {
				open -= 1
			}
		}
		first := braceExpansionRecII(e[1:i])
		second := braceExpansionRecII(e[i+1:])
		return concatenate(first, second)
	} else {
		sb := strings.Builder{}
		var i int
		for i = 0; i < n && unicode.IsLetter(rune(e[i])); i++ {
			sb.WriteByte(e[i])
		}
		return concatenate([]string{sb.String()}, braceExpansionRecII(e[i:]))
	}
}

func braceExpansionBFSII(expression string) []string {
	resultMap := make(map[string]bool)
	q := []string{expression}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		left := strings.Index(cur, OPENING_BRACE)
		if left == -1 {
			resultMap[cur] = true
			continue
		}
		i := left
		for i < len(cur) && string(cur[i]) != CLOSING_BRACE {
			if string(cur[i]) == OPENING_BRACE {
				left = i
			}
			i++
		}
		right := i
		processed := cur[:left]
		processing := strings.Split(cur[left+1:right], ",")
		unprocessed := cur[right+1:]

		for _, part := range processing {
			sb := strings.Builder{}
			sb.WriteString(processed)
			sb.WriteString(part)
			sb.WriteString(unprocessed)
			q = append(q, sb.String())
		}
	}

	result := []string{}
	for k := range resultMap {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}

func BraceExpansionBFSII(expression string) []string {
	return braceExpansionBFSII(expression)
}

func BraceExpansion´RecII(expression string) []string {
	return braceExpansionRecII(expression)
}
