// @algorithm @lc id=1711 lang=golang
// @title find-valid-matrix-given-row-and-column-sums

package main

// @test([3,8],[4,7])=[[3,0],[1,7]]
// @test([5,7,10],[8,6,8])=[[0,5,0],[6,1,0],[2,0,8]]
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	result := make([][]int, len(rowSum))
	for i := range result {
		result[i] = make([]int, len(colSum))
	}
	for i, j := 0, 0; i < len(rowSum) && j < len(colSum); {
		if rowSum[i] < colSum[j] {
			result[i][j] = rowSum[i]
			colSum[j] -= rowSum[i]
			rowSum[i] = 0
			i++
		} else if rowSum[i] > colSum[j] {
			result[i][j] = colSum[j]
			rowSum[i] -= colSum[j]
			colSum[j] = 0
			j++
		} else {
			result[i][j] = colSum[j]
			rowSum[i] = 0
			colSum[j] = 0
			i++
			j++
		}
	}
	return result
}
