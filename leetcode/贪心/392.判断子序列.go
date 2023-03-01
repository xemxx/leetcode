/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	si := 0
	for i := 0; i < len(t); i++ {
		if s[si] == t[i] {
			si++
			if si == len(s) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

