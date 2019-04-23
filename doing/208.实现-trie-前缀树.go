package main

type Trie struct {
	count int
	child map[string]Trie
}

func main() {
	word := "dasd"
	prefix := "da"
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search(word)
	param_3 := obj.StartsWith(prefix)
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{0, map[string]Trie{}}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	tmp:=
	for i := 0; i < len(word); i++ {
		v := word[i : i+1]
		_, ok := t.child[v]
		if ok {
			continue
		}
		t.count += 1
		t.child
	}
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (57.24%)
 * Total Accepted:    5.6K
 * Total Submissions: 9.7K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
 */
