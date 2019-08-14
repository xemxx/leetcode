## 方法一：暴力法

直接从左上角开始，依次两个层循环即可

## 方法二：二分查找

分为两种：

- 对每一行进行二分查找，同时可以在二分之前，直接判断第一个和最后一个元素与目标值的大小，可以缩减部分不需要查找的行
- 遍历对角线的元素，通过每一个遍历的元素对其所在行和列进行二分查找

## 方法三

说明：

- 本方法属于巧妙解法，主要是因为本矩阵的每一行以及每一列都是升序的
- 从右上角的元素`(row,col)`开始，如果当前元素大于目标值，说明需要找更小的，则`col--`，如果当前元素小于目标值，说明需要找更大的，则`row++`
- 或者从左下角开始，只是相反处理`row`和`col`即可
- 不能从左上角和右下角开始的原因是因为两个方向都是大于当前值或者小于当前值，移动可能丢失一行的可能值或者一列的可能值

代码：
``` golang
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row <= len(matrix)-1 && col >= 0 {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return true
		}
	}
	return false
}
```
