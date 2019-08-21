package main

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	arr := make([]int, 0, len(nums))
	arr = append(arr, nums[0])
	for _, v := range nums[1:] {
		if v > arr[len(arr)-1] {
			arr = append(arr, v)
		} else {
			l, r := 0, len(arr)-1
			for l < r {
				mid := (l + r) / 2
				if arr[mid] >= v {
					r = mid
				} else {
					l = mid + 1
				}
			}
			arr[l] = v
		}
	}
	return len(arr)
}
