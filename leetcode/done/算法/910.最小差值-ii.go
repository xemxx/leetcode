package LEETCODE

/*
 * @lc app=leetcode.cn id=910 lang=golang
 *
 * [910] 最小差值 II
 */

// @lc code=start
import "sort"

func smallestRangeII(A []int, K int) int {
	sort.Sort(sort.IntSlice(A))
	result := A[len(A)-1] - A[0] //特殊情况，要么全部上移要么全部下移
	// 寻找切切割点
	for i := 0; i < len(A)-1; i++ {
		max := Max(A[i]+K, A[len(A)-1]-K)
		min := Min(A[0]+K, A[i+1]-K)
		tmp := max - min
		result = Min(tmp, result)
	}
	return result
}
func Max(a, b int) int {
	if b > a {
		a = b
	}
	return a
}
func Min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

// @lc code=end
