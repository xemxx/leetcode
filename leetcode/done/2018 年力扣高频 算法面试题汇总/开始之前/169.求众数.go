package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))

}

// 暴力
func majorityElement(nums []int) int {

	for _, i := range nums {
		a := 0
		for _, j := range nums {
			if i == j {
				a++
			}
		}
		if a > len(nums)/2 {
			return i
			break
		}
	}
	return 0
}

//思路： 1、hash计算次数，2、排序取中值，3、如下方法
func majorityElement1(nums []int) int {
	count := 0
	x := nums[0]
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			x = nums[i]
		}
		if nums[i] == x {
			count++
		} else {
			count--
		}
	}
	return x
}

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 求众数
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (58.48%)
 * Total Accepted:    35.6K
 * Total Submissions: 60.8K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在众数。
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */
