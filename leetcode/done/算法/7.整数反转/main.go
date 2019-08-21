/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-100101))
}

func reverse(x int) int {
	re := 0
	for x != 0 {
		re += x % 10
		re *= 10
		x = x / 10
	}
	re /= 10
	if re > math.MaxInt32 || re < math.MinInt32 {
		return 0
	}
	return re

}
