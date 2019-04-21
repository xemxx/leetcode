package main

import (
	"fmt"
	"runtime"
)

var result []string

func main() {
	runtime.GOMAXPROCS(4)
	//深搜日常超时，等待下一个解决办法
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak1(s, wordDict))

}

//第一种，超时
func wordBreak(s string, wordDict []string) []string {
	result = []string{}
	cut := make([]bool, len(s)+1)
	dfs(s, wordDict, 0, cut)
	return result
}

func dfs(s string, wordDict []string, index int, cut []bool) {
	l := len(s)
	if index == l {
		res := ""
		h := 0
		for i := 0; i <= l; i++ {
			if cut[i] {
				res += s[h:i] + " "
				h = i
			}
		}
		res = res[0 : len(res)-1]
		result = append(result, res)
	}
	for i := index + 1; i <= l; i++ {
		word := s[index:i]
		if inword(word, wordDict) {
			cut[i] = true
			fmt.Println(i)
			dfs(s, wordDict, i, cut)
			cut[i] = false
		}
	}
}

func inword(word string, wordDict []string) bool {
	for _, v := range wordDict {
		if v == word {
			return true
		}
	}
	return false
}

//第二种，先动规，再深搜
func wordBreak1(s string, wordDict []string) []string {
	result = []string{}
	//cut := make([]bool, len(s)+1)
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
	dfs1(s, wordDict, 0, f, "")
	return result
}

func dfs1(s string, wordDict []string, index int, f []bool, res string) {
	l := len(s)
	if index == l {
		res = res[0 : len(res)-1]
		fmt.Println(res)
		result = append(result, res)
	}
	for i := index + 1; i <= l; i++ {
		word := s[index:i]
		if inword(word, wordDict) {
			res += word + " "
			dfs1(s, wordDict, i, f, res)
		}
	}
}

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 *
 * https://leetcode-cn.com/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (36.10%)
 * Total Accepted:    2.9K
 * Total Submissions: 8K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典
 * wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
 *
 * 说明：
 *
 *
 * 分隔时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入:
 * s = "catsanddog"
 * wordDict = ["cat", "cats", "and", "sand", "dog"]
 * 输出:
 * [
 * "cats and dog",
 * "cat sand dog"
 * ]
 *
 *
 * 示例 2：
 *
 * 输入:
 * s = "pineapplepenapple"
 * wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
 * 输出:
 * [
 * "pine apple pen apple",
 * "pineapple pen apple",
 * "pine applepen apple"
 * ]
 * 解释: 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入:
 * s = "catsandog"
 * wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出:
 * []
 *
 *
 */
