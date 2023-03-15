package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(beautifulSubarrays([]int{4, 3, 1, 2, 4}))
	fmt.Println(findMinimumTime([][]int{{1, 18, 5}, {3, 15, 1}}))
}

func vowelStrings(words []string, left int, right int) int {
	m := map[byte]struct{}{'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{}}
	count := 0
	for i := left; i <= right; i++ {
		if _, ok := m[words[i][0]]; ok {
			if _, ok := m[words[i][len(words[i])-1]]; ok {
				count++
			}
		}
	}
	return count
}

func maxScore(nums []int) int {
	// 实际只需要整体排序即可
	arr2 := make([]int, 0, len(nums))
	cnt := 0
	arr1 := 0
	for _, v := range nums {
		if v > 0 {
			arr1 += v
			cnt++
		} else if v == 0 {
			cnt++
		} else {
			arr2 = append(arr2, v)
		}
	}
	result := 0
	sort.Ints(arr2)
	for i := len(arr2) - 1; i >= 0; i-- {
		if arr1-arr2[i] > 0 {
			result++
			arr1 -= arr2[i]
		} else {
			break
		}
	}
	return result + cnt
}

func beautifulSubarrays(nums []int) int64 {
	// 每次选两个数值去掉k，相当于两个数的 同一个 二进制位做异或最终等于0
	// 所以子数组最终为0，代表所有的数的异或和最终为0
	// 按照前缀和的思想，假定abc是可以成功的组合，那么 x^a^b^c == x ，即算出所有的异或和，当有两个异或和相等，则这中间的子数组满足条件，再有多余的相等值，则可以组合成多个新的子数组
	// 这里需要提前安排一个 s:1 记作nums[0]的异或和出现过一次，在下一次碰到结果一致时，能直接加1，否则会漏一次
	s := 0
	m := map[int]int{0: 1}
	result := 0
	for _, v := range nums {
		s ^= v
		result += m[s]
		fmt.Println(result, s)
		m[s]++
	}
	return int64(result)
}

func findMinimumTime(tasks [][]int) int {
	run := make([]bool, 2001)
	// 排序后贪心，每次先分配到后面，然后在下一个数组内把前面已经分配过的使用一遍，然后把剩下的再放到后面
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	for _, v := range tasks {
		s, e, d := v[0], v[1], v[2]
		for _, vv := range run[s : e+1] {
			if vv {
				d--
			}
		}
		for j := 0; j < e && d > 0; j++ {
			if run[e-j] {
				continue
			}
			run[e-j] = true
			d--
		}
	}
	result := 0
	for _, v := range run {
		if v {
			result++
		}
	}
	return result
}
