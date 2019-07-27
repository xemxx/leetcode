package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

func twoSum(nums []int, target int) []int {
	re := make(map[int]int, len(nums))
	for k, v := range nums {
		now := target - v
		_, ok := re[now]
		if ok {
			return []int{re[now], k}
		}
		re[v] = k
	}
	return []int{}
}

// 如果给定为顺序数组，可使用双指针一次遍历即可实现
