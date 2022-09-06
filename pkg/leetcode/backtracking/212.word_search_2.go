package backtracking

/*
212. Word Search II
https://leetcode.com/problems/word-search-ii/
*/
type Trie struct {
    children map[byte]*Trie
    word     string
}

func (t *Trie) Insert(word string) {
    node := t
    for i := range word {
        ch := word[i]
        if node.children[ch] == nil {
            node.children[ch] = &Trie{children: map[byte]*Trie{}}
        }
        node = node.children[ch]
    }
    node.word = word
}

func findWords(board [][]byte, words []string) (ans []string) {
	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    t := &Trie{children: map[byte]*Trie{}}
    for _, word := range words {
        t.Insert(word)
    }

    m, n := len(board), len(board[0])

    var floodFill func(node *Trie, x, y int)
    floodFill = func(node *Trie, x, y int) {
        ch := board[x][y]
        nxt := node.children[ch]
        if nxt == nil {
            return
        }

        if nxt.word != "" {
            ans = append(ans, nxt.word)
            nxt.word = ""
        }

        if len(nxt.children) > 0 {
            board[x][y] = '#'
            for _, d := range dirs {
                nx, ny := x+d.x, y+d.y
                if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
                    floodFill(nxt, nx, ny)
                }
            }
            board[x][y] = ch
        }
        
        if len(nxt.children) == 0 {
            delete(node.children, ch)
        }
    }
    for i, row := range board {
        for j := range row {
            floodFill(t, i, j)
        }
    }

    return
}


func FindWords(board [][]byte, words []string) []string {
	return findWords(board, words)
}