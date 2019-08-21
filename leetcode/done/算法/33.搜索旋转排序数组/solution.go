package main

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		// 加此判断是为了可以提前结束查找
		if nums[mid] == target {
			return mid
		}

		// 二分搜索的具体实现，每次去除一半的区域
		if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

// 可以采用异或版本题解：https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/ji-jian-solution-by-lukelee/
// golang的异或运算符不支持bool类型因此不做示例
