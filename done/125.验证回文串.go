package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	if isPalindrome(s) {
		fmt.Println("1")
	}

}

func isPalindrome(s string) bool {
	l := len(s)
	j := l - 1
	i := 0
	for i < j {
		a := rune(s[i])
		b := rune(s[j])
		if !unicode.IsLetter(a) && !unicode.IsDigit(a) {
			i++
		} else if !unicode.IsLetter(b) && !unicode.IsDigit(b) {
			j--
		} else {
			if !strings.EqualFold(string(s[i]), string(s[j])) {
				return false
			}
			i++
			j--
		}
	}
	return true
}

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (38.94%)
 * Total Accepted:    31.8K
 * Total Submissions: 81.6K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */
