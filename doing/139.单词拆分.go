package main

import "fmt"

var result [][]string

func main() {
	//很简单的想到深搜，结果一下案例超时，我自己的mac本地跑了500多秒没跑完。。。下面用动规试试
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak(s, wordDict))

}
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
