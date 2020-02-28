## 题目

给定一个k和一个整数数组，每个数组元素都可+k或者-k得到新的数组元素，但每个数只能变换一次，在无数的可能数组中，求最大值和最小值的差值的最小值

## 错误思路

每次保存到目前为止的最小值，并保存求的此最小值的最大值和最小值，每次有一个新的数，分别+k和-k，求得最小值，更新之前保存的值，最后遍历完即可得

没有ac，我也不知道为啥，，，暂时没有更好的思路

## 题解

[LeetCode原题910](https://leetcode-cn.com/problems/smallest-range-ii/)

``` go
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
```