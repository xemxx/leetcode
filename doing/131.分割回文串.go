package main

var result [][]string
var cut [][]bool

func main() {
	s := "3123"
	partition(s)

}

func partition(s string) [][]string {
	find(s, 0)
	return result
}

func find(s string, index int) {
	if index == len(s) {

	}
	for i := index; index < len(s); i++ {
		if isPa(s, index, i) {
			find(s, i+1)
			cut[index][i] = true
		}
	}
}
func isPa(s string, h int, l int) bool {
	for h < l && s[h] == s[l] {
		h++
		l--
	}
	return h >= l
}

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (61.88%)
 * Total Accepted:    5.9K
 * Total Submissions: 9.6K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 *
 * 返回 s 所有可能的分割方案。
 *
 * 示例:
 *
 * 输入: "aab"
 * 输出:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 */
