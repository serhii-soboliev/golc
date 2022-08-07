package backtracking

/*
212. Word Search II
https://leetcode.com/problems/word-search-ii/
*/
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	curr := t
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	 curr := t
	 for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil{
			return false
		}
		curr = curr.children[idx]
	 }
	 return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t
	for _, ch := range prefix {
	   idx := ch - 'a'
	   if curr.children[idx] == nil{
		   return false
	   }
	   curr = curr.children[idx]
	}
	return true
}

func (t *Trie) GetSubTrie(r rune) Trie {
	idx := r - 'a'	
	return *t.children[idx]
}

func findWords(board [][]byte, words []string) []string {

	n, m := len(board), len(board[0])

	onBoard := func(i, j int) bool {
		return i >= 0 && i < n && j >=0 && j < m
	}

	resWithDuplicates := [] string {}
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	
	directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var dfs func(subTrie Trie, i int, j int, visited [][]bool, prefix string) []string

	dfs = func(subTrie Trie, i int, j int, visited [][]bool, prefix string) []string {
		subresult := []string{}
		if subTrie.isEnd {
			subresult = []string{prefix}
			subTrie.isEnd = false
		}
		visited[i][j] = true

		for _, dir := range directions {
			nI := i + dir[0]
			nJ := j + dir[1]

			if onBoard(nI, nJ) && !visited[nI][nJ] {
				b := board[nI][nJ]
				s := string(b)
			    if subTrie.StartsWith(s) {
					subresult = append(subresult, 
						dfs(subTrie.GetSubTrie(rune(b)), nI, nJ, visited, prefix+s )...)
				}
			}
		}
		visited[i][j] = false
		return subresult
	}

	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			b := board[i][j]
			s := string(b)

			if !trie.StartsWith(s) { continue }

			visited := make([][]bool, n)
			for i := range visited {
				visited[i] = make([]bool, m)
			}

			resWithDuplicates = append(resWithDuplicates, dfs(trie.GetSubTrie(rune(b)), i, j, visited, s) ...)
		}
	}

	resMap := make(map[string]bool) 
	for _, w := range resWithDuplicates {
		resMap[w] = true
	}
	result := []string{}
	for k := range resMap {
		result = append(result, k)
	}
	return result
}

func FindWords(board [][]byte, words []string) []string {
	return findWords(board, words)
}