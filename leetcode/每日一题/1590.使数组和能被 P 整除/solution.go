// @algorithm @lc id=1694 lang=golang
// @title make-sum-divisible-by-p

package main

// @test([3,1,4,2],6)=1
// @test([6,3,5,2],9)=2
// @test([1,2,3],3)=0
// @test([26,19,11,14,18,4,7,1,30,23,19,8,10,6,26,3], 26)=3
// @test([8,32,31,18,34,20,21,13,1,27,23,22,11,15,30,4,2], 148)=7
func minSubarray(nums []int, p int) int {
	// s 前缀和
	// s[right]−s[left]≡x(mod p)
	// s[right]−x≡s[left](mod p)
	if len(nums)%p == 0 {
		return 0
	}
	x := 0
	for _, v := range nums {
		x += v
	}
	if x < p {
		return -1
	}
	x %= p
	if x == 0 {
		return 0
	}
	s := 0
	result := len(nums)
	// 先提供一个基本的0，但是因为i从0开始，有可能一路相加然后得到结果，所以需要取-1
	last := map[int]int{0: -1}
	for i, v := range nums {
		s += v
		last[s%p] = i
		if j, ok := last[(s-x+p)%p]; ok {
			if i-j < result {
				result = i - j
			}
		}
	}
	if result == len(nums) {
		return -1
	}
	return result
}
