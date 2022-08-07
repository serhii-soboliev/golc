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

func findWords(board [][]byte, words []string) []string {

	contains := func(i int, j int, a [][]int) bool {
		for _, v := range a {
			if v[0] == i && v[1] == j {
				return true
			}
		}
		return false
	}
	n, m := len(board), len(board[0])

	onBoard := func(i, j int) bool {
		return i >= 0 && i < n && j >=0 && j < m
	}

	result := [] string {}
	trie := Constructor()
	
	directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var bfs func(i int, j int, visited [][]int, prefix string)

	bfs = func(i int, j int, visited [][]int, prefix string) {
		trie.Insert(prefix)
		if contains(i, j, visited) {	
			return
		}
		b := board[i][j]
		newPrefix := prefix + string(rune(b))
		trie.Insert(newPrefix)
		for _, dir := range directions {
			newI := i + dir[0]
			newJ := j + dir[1]
			if onBoard(newI, newJ) {
				visited = append(visited, []int{i, j})
				bfs(newI, newJ, visited, newPrefix)
				visited = visited[:len(visited) - 1]
			}
		}
	}

	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			bfs(i, j, [][]int{}, "")
		}
	}

	for _, word := range words {
		if trie.Search(word) {
			result = append(result, word)
		}
	}

	return result
}

func FindWords(board [][]byte, words []string) []string {
	return findWords(board, words)
}