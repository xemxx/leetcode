package main

import "fmt"

var result [][]string

func main() {
	//很简单的想到深搜，结果一下案例超时，我自己的mac本地跑了500多秒没跑完。。。下面用动规试试
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak3(s, wordDict))

}

//第一种递归，遍历字典
func wordBreak(s string, wordDict []string) bool {
	return dfs(s, wordDict, 0)
}

func dfs(s string, wordDict []string, index int) bool {

	if index == len(s) {
		return true
	}
	now := s[index:]
	for _, v := range wordDict {
		l := len(v)
		sl := len(now)
		//sl与l的判断是为了防止now[0:l]越界报错
		if sl >= l && now[0] == v[0] && now[0:l] == v {
			if dfs(s, wordDict, index+l) {
				//只要找到一个成功的就直接成功
				return true
			}
		}
	}
	return false
}

//第二种递归，遍历字符串
func wordBreak2(s string, wordDict []string) bool {
	return dfs2(s, wordDict, 0)
}

func dfs2(s string, wordDict []string, index int) bool {

	if index == len(s) {
		return true
	}
	for i := index + 1; i <= len(s); i++ {
		fmt.Println(i)
		if inword(s[index:i], wordDict) && dfs2(s, wordDict, i) {
			//只要找到一个成功的就直接成功
			return true
		}
	}
	return false
}

func inword(s string, wordDict []string) bool {
	for _, v := range wordDict {
		if s == v {
			return true
		}
	}
	return false
}

//第二种优化 成功通过
func wordBreak2_1(s string, wordDict []string) bool {
	//return dfs2(s, wordDict, 0)
	memo := make([]bool, len(s)+1)
	return dfs2_1(s, wordDict, 0, memo)
}
func dfs2_1(s string, wordDict []string, index int, memo []bool) bool {

	if index == len(s) {
		return true
	}
	if memo[index] {
		return false
	}
	for i := index + 1; i <= len(s); i++ {
		if memo[i] {
			continue
		}
		//此处判断是否在字典内可以优化为map[string]bool 直接判断
		if inword(s[index:i], wordDict) && dfs2_1(s, wordDict, i, memo) {
			//只要找到一个成功的就直接成功
			return true
		}
	}
	//记录已经查过的情况，因为slice是引用类型
	memo[index] = true
	return false
}

//第三种动态规划 能过leetcode f[i]表示i长度的字符串能否被字典拼接
func wordBreak3(s string, wordDict []string) bool {
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 1; i <= len(s); i++ {
		for _, v := range wordDict {
			if i >= len(v) && f[i-len(v)] && v == s[i-len(v):i] {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode-cn.com/problems/word-break/description/
 *
 * algorithms
 * Medium (39.89%)
 * Total Accepted:    7.7K
 * Total Submissions: 19.3K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 *
 * 说明：
 *
 *
 * 拆分时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
 *
 *
 * 示例 2：
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
 * 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 */
