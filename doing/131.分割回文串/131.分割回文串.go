package main

import "fmt"

var result [][]string

func main() {
	s := "aab"
	fmt.Println(partition(s))

}

func partition(s string) [][]string {
	cut := make([]bool, len(s))
	//消除在leetcode的多个测试案例中全局变量的影响
	result = [][]string{}
	dfs(s, 0, cut)
	return result
}

func dfs(s string, index int, cut []bool) {
	l := len(s)
	if index == l {
		//搜索到底即可进行分割保存此次方案
		now := []string{}
		for i, j := 0, 0; i < l; i++ {
			if cut[i] {
				now = append(now, s[j:i+1])
				j = i + 1
			}
		}
		result = append(result, now)
	}
	for i := index; i < l; i++ {
		if isPa(s, index, i) {
			//先标记
			cut[i] = true
			//进入下一层
			dfs(s, i+1, cut)
			//取消标记，准备下次深搜
			cut[i] = false
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
