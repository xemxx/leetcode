// @algorithm @lc id=1000033 lang=golang
// @title find-longest-subarray-lcci

package main

// @test(["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"])=["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
// @test(["A","A"])=[]
func findLongestSubarray(array []string) []string {
	// 前缀和
	chartCnt := 0
	numCnt := 0
	m := map[int]int{0: -1}
	r, l := -1, -1
	result := 0
	for i, v := range array {
		if v[0] < '0' || v[0] > '9' {
			chartCnt++
		} else {
			numCnt++
		}
		v, ok := m[chartCnt-numCnt]
		if ok {
			if i-v > result {
				r = v
				l = i
				result = i - v
			}
		} else {
			m[chartCnt-numCnt] = i
		}
	}
	return array[r+1 : l+1]
}
