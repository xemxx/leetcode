package main

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
func maxSubArray(nums []int) int {
	max, iMax := nums[0], nums[0]
	for _, v := range nums[1:] {
		tmp := v + iMax
		if tmp < v {
			tmp = v
		}
		iMax = tmp
		if tmp > max {
			max = tmp
		}
	}
	return max
}
