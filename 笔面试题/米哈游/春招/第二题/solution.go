package main

import "fmt"

/**
 *
 * @param s string字符串
 * @return string字符串
 */
func minRemove(s string) string {
	// write code here
	left, right := 0, 0
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		}
	}
	if left <= right {
		count := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				count++
				result += "("
			} else if s[i] == ')' {
				if count > 0 {
					result += ")"
					count--
				}
			} else {
				result = fmt.Sprintf("%s%c", result, s[i])
			}
		}
	} else if left > right {
		count := 0
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == ')' {
				count++
				result += ")"
			} else if s[i] == '(' {
				if count > 0 {
					result += "("
					count--
				}
			} else {
				result = fmt.Sprintf("%s%c", result, s[i])
			}
		}
		tmp := ""
		for i := len(result) - 1; i >= 0; i-- {
			tmp = fmt.Sprintf("%s%c", tmp, result[i])
		}
	}
	return result
}
