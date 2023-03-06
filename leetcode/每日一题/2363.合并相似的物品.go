/*
 * @lc app=leetcode.cn id=2363 lang=golang
 *
 * [2363] 合并相似的物品
 */
package leetcode

// @lc code=start
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	keyM := make([]int, 1001)
	// valM := make(map[int]int)
	for i := range keyM {
		keyM[i] = -1
	}
	for _, v := range items1 {
		keyM[v[0]] = 0
		keyM[v[0]] = v[1]
	}
	for _, v := range items2 {
		if keyM[v[0]] == -1 {
			keyM[v[0]] = 0
		}
		keyM[v[0]] += v[1]
	}
	maxLen := len(items1)
	if len(items2) > maxLen {
		maxLen = len(items2)
	}
	v := make([][]int, 0, maxLen)
	for kk, vv := range keyM {
		if vv > 0 {
			v = append(v, []int{kk, keyM[kk]})
		}
	}
	return v
}

// @lc code=end
