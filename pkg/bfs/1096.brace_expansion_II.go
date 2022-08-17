package bfs

import (
	"strings"
	"sort"
)

/*
1096. Brace Expansion II
https://leetcode.com/problems/brace-expansion-ii/
*/
var OPENING_BRACE = "{"
var CLOSING_BRACE = "}"

func braceExpansionII(expression string) []string {
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
		i:=left; 
		for i < len(cur) && string(cur[i]) != CLOSING_BRACE {
			if string(cur[i]) == OPENING_BRACE { left = i }
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

func BraceExpansionII(expression string) []string {
	return braceExpansionII(expression)
}