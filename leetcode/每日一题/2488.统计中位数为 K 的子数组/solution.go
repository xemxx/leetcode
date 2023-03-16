// @algorithm @lc id=2574 lang=golang
// @title count-subarrays-with-median-k

package main

// @test([3,2,1,4,5],4)=3
// @test([2,3,1],3)=1
func countSubarrays(nums []int, k int) int {
	// 分别前后做前缀和，然后两两做乘，最多遍历 2n
	left := map[int]int{0: 0}
	right := map[int]int{0: 0}
	result := 0
	for idx, v := range nums {
		if v == k {
			tmp := 0
			for i := idx - 1; i >= 0; i-- {
				if nums[i] < k {
					tmp--
				} else {
					tmp++
				}
				left[tmp]++
			}
			tmp = 0
			for i := idx + 1; i < len(nums); i++ {
				if nums[i] < k {
					tmp--
				} else {
					tmp++
				}
				right[tmp]++
			}
			result += right[1] + right[0] + left[1] + left[0]
			for i := range right {
				result += right[i]*left[0-i] + right[i]*left[0-i+1]
			}
			break
		}
	}
	return result + 1
}
