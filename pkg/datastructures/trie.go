package datastructures

/*
208. Implement Trie (Prefix Tree)
https://leetcode.com/problems/implement-trie-prefix-tree/
*/
type Trie struct {
	children [26]*Trie
    isEnd bool;
}

func Constructor() Trie {
	return Trie{}
}


func (this *Trie) Insert(word string)  {
	
}


func (this *Trie) Search(word string) bool {
	return false
}


func (this *Trie) StartsWith(prefix string) bool {
	return false
}