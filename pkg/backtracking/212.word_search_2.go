package backtracking

/*
212. Word Search II
https://leetcode.com/problems/word-search-ii/
*/
type Trie struct {
	children [26]*Trie
	isEnd    bool
	word 	 string
}

func buildTrie(words []string) *Trie {
    root := &Trie{}
    for _, word := range words {
        tmp := root
        for _, c := range word {
            i := c - 'a'
            if tmp.children[i] == nil { tmp.children[i] = &Trie{} }
            tmp = tmp.children[i]
        }
        tmp.isEnd = true
		tmp.word = word
    }
    return root
}

func findWords(board [][]byte, words []string) []string {
	root := buildTrie(words)
	result := []string{}

	var floodFill func(board [][]byte, root *Trie, i, j int) 

	floodFill = func(board [][]byte, root *Trie, i, j int) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) { return }
		tmp := board[i][j]
		if tmp == '$' || root.children[tmp - 'a'] == nil { return }
		root = root.children[tmp - 'a']
		if root.isEnd { 
			result = append(result, root.word); 
			root.word = "" 
			root.isEnd = false
		}
		board[i][j] = '$'
		floodFill(board, root, i + 1, j)
		floodFill(board, root, i - 1, j)
		floodFill(board, root, i, j + 1)
		floodFill(board, root, i, j - 1)
		board[i][j] = tmp
	}

	n,m := len(board), len(board[0])
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			floodFill(board, root, i, j)
		}
	}
	
	return result
	
}

func FindWords(board [][]byte, words []string) []string {
	return findWords(board, words)
}