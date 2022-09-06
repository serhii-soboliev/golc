package trie

/*
208. Implement Trie (Prefix Tree)
https://leetcode.com/problems/implement-trie-prefix-tree/
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
