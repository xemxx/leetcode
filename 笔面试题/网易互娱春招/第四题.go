package main

import "sort"

/**
 *
 * @param boxes int整型二维数组
 * @return int整型
 */
func maxBoxes(boxes [][]int) int {
	// write code here
	sort.Sort(SortBy(boxes))
	max := 1
	cur := 0
	for i := 1; i < len(boxes); i++ {
		if boxes[i][0] > boxes[cur][0] && boxes[i][1] > boxes[cur][1] && boxes[i][2] > boxes[cur][2] {
			cur = i
			max++
		}
	}
	return max
}

type SortBy [][]int

func (a SortBy) Len() int {
	return len(a)
}
func (a SortBy) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortBy) Less(i, j int) bool {
	return a[i][0]*a[i][1]*a[i][2] < a[j][0]*a[j][1]*a[j][2]
}
