package main

/**
 * 返回两个数组中相同数字不相交的最大连线数
 * @param A int整型一维数组 整数集合A
 * @param B int整型一维数组 整数集合B
 * @return int整型
 */
func maxUncrossedLines(A []int, B []int) int {
	// write code here
	ma := map[int]int{}
	mb := map[int]int{}
	count := 0
	la, lb := len(A), len(B)
	i, j := 0, 0
	for i < la && j < lb {
		a, b := A[i], B[j]
		ma[a], mb[b] = i, j
		if v, ok := ma[b]; ok {
			count++
			i = v
			ma = map[int]int{}
			mb = map[int]int{}
		} else if v, ok := mb[a]; ok {
			count++
			j = v
			mb = map[int]int{}
			ma = map[int]int{}
		}
		i++
		j++

	}
	return count
}
