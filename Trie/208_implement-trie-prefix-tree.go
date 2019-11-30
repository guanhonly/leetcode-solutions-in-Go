/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/implement-trie-prefix-tree/
 */

package Trie

type TrieNode struct {
	val      int32
	children [26]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			val: ' ',
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	ws := this.root
	for _, c := range word {
		if ws.children[c-'a'] == nil {
			ws.children[c-'a'] = &TrieNode{
				val: c,
			}
		}
		ws = ws.children[c-'a']
	}
	ws.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	ws := this.root
	for _, c := range word {
		if ws.children[c-'a'] == nil {
			return false
		}
		ws = ws.children[c-'a']
	}
	return ws.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	ws := this.root
	for _, c := range prefix {
		if ws.children[c-'a'] == nil {
			return false
		}
		ws = ws.children[c-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
