/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */
package leetcode

import "fmt"

// @lc code=start
func validPalindrome(s string) bool {
	if is(s) {
		return true
	}
	for i := 0; i < len(s); i++ {
		tmp := s[:i] + s[i+1:]
		fmt.Println(tmp)
		if is(tmp) {
			return true
		}
	}
	return false
}

func is(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end
