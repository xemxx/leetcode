package main

/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 */
func longestSubstring(s string, k int) int {
	l := len(s)
	if l == 0 || l < k {
		return 0
	}
	if k < 2 {
		return l
	}
	return find(s, k)
}

func find(s string, k int) int {
	l, r := 0, len(s)-1
	if r-l+1 < k {
		return 0
	}
	m := map[byte]int{}
	for _, k := range s {
		m[byte(k)]++
	}
	for m[s[l]] < k && l > r {
		l++
	}
	for m[s[r]] < k && l < r {
		r--
	}
	if r-l+1 < k {
		return 0
	}
	for i := l; i <= r; i++ {
		if m[s[i]] < k {
			left := find(s[l:i], k)
			right := find(s[i+1:r+1], k)
			if left < right {
				left = right
			}
			return left
		}
	}
	return r - l + 1
}
