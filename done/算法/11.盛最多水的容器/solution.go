package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		now := 0
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
