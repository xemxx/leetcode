package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

//双指针解法 时间复杂度为log（n）
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		now := 0
		// 每次遍历将最短边更换
		if height[i] < height[j] {
			now = height[i] * (j - i)
			i++
		} else {
			now = height[j] * (j - i)
			j--
		}
		if max < now {
			max = now
		}
	}
	return max
}
