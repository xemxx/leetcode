package main

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */
func longestConsecutive(nums []int) int {
	mLen := 0
	mmp := make(map[int]int, len(nums))
	for _, k := range nums {
		_, ok := mmp[k]
		if ok == false {
			left := mmp[k-1]
			right := mmp[k+1]

			curLen := 1 + left + right

			if curLen > mLen {
				mLen = curLen
			}
			mmp[k] = curLen
			mmp[k-left] = curLen
			mmp[k+right] = curLen
		}
	}
	return mLen
}
