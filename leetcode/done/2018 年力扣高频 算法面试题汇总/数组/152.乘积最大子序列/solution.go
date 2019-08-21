package main

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */
func maxProduct(nums []int) int {
	iMax, iMin, Max := nums[0], nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = max(iMax*v, v)
		iMin = min(iMin*v, v)
		Max = max(Max, iMax)
	}
	return Max
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}
