package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))

}

func singleNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		a := 2
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {
				a = 1
				break
			}
		}
		if a == 2 {
			return nums[i]
			break
		}
	}
	return 0
}

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 *
 * https://leetcode-cn.com/problems/single-number/description/
 *
 * algorithms
 * Easy (60.25%)
 * Total Accepted:    59.4K
 * Total Submissions: 98.6K
 * Testcase Example:  '[2,2,1]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *
 * 说明：
 *
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * 示例 1:
 *
 * 输入: [2,2,1]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入: [4,1,2,1,2]
 * 输出: 4
 *
 */
