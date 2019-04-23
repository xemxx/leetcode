package main

import "fmt"

func main() {
	word := "apple"
	prefix := "apple"
	obj := Constructor()
	obj.Insert(word)
	fmt.Println(obj.Search(prefix))
	fmt.Println(obj.StartsWith(prefix))
}

// type Trie struct {
// 	count int
// 	child map[string]*Trie
// }
// func Constructor() Trie {
// 	return Trie{0, map[string]*Trie{}}
// }

// func (t *Trie) Insert(word string) {
// 	node := t
// 	for i := 0; i < len(word); i++ {
// 		v := word[i : i+1]
// 		_, ok := node.child[v]
// 		if !ok {
// 			node.count++
// 			newNode := Trie{0, map[string]*Trie{}}
// 			node.child[v] = &newNode
// 		}
// 		node = node.child[v]
// 	}
// }

// func (t *Trie) Search(word string) bool {
// 	node := t
// 	previous := false
// 	for i := 0; i < len(word); i++ {
// 		v := word[i : i+1]
// 		_, ok := node.child[string(v)]
// 		if ok {
// 			//fmt.Println(3, v)
// 			previous = true
// 			node = node.child[string(v)]
// 		} else if previous == true {
// 			//fmt.Println(2)
// 			return false
// 		} else {
// 			//fmt.Println(1)
// 			for _, v := range node.child {
// 				if v.Search(word) {
// 					return true
// 				}
// 			}
// 			//此处应终止外层循环，防止产生断层匹配
// 			return false
// 		}
// 	}
// 	return previous
// }

// func (t *Trie) StartsWith(prefix string) bool {
// 	node := t
// 	for _, v := range prefix {
// 		_, ok := node.child[string(v)]
// 		if ok {
// 			node = node.child[string(v)]
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

//由于上述代码理解错误tire的意思，作废，我也不知道我写了个什么数据结构出来

type Trie struct {
	count int //用于统计词频
	isKey bool
	child map[string]*Trie //此处可用[26]*Tire继续优化，但考虑到实际业务，用map更为稳妥
}

func Constructor() Trie {
	return Trie{0, false, map[string]*Trie{}}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		v := word[i : i+1]
		_, ok := node.child[v]
		if !ok {
			newNode := Trie{0, false, map[string]*Trie{}}
			node.child[v] = &newNode
		}
		node = node.child[v]
	}
	node.count++
	node.isKey = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, v := range word {
		node = node.child[string(v)]
		if node == nil {
			return false
		}
	}
	return node.isKey
	// 对性能的强迫症使我用了上述代码
	// for i := 0; i < len(word) && node != nil; i++ {
	// 	v := word[i : i+1]
	// 	node = node.child[v]
	// }
	// if node != nil && node.isKey {
	// 	return true
	// }
	// return false
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, v := range prefix {
		node = node.child[string(v)]
		if node == nil {
			return false
		}
	}
	return true
	// 同上
	// for _, v := range prefix {
	// 	_, ok := node.child[string(v)]
	// 	if ok {
	// 		node = node.child[string(v)]
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true
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
