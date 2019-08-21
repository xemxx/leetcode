## 动态规划
核心思想：用一个数组`int arr[]`辅助，`arr[i]`代表到`i`为止的最长上升子序列的长度，然后由公式`arr[i]=max(arr[j])+1`得到`arr[i]`的值，前提是`j<i`
具体实现：
- 每次遍历`nums[i]`，开始遍历`arr[]`，选出最大的值然后再`+1`得到最后`arr[i]`的值即可。

代码：
``` golang
func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	max := 0
	arr := make([]int, len(nums))
	for i := range nums {
		tmp := 0
		for j := range arr {
			if nums[j] < nums[i] {
				if arr[j] > tmp {
					tmp = arr[j]
				}
			}
		}
		arr[i] = tmp + 1
		if arr[i] > max {
			max = arr[i] + 1
		}
	}
	return max
}
```

时间复杂度为O(n^2)

## 动归+二分
原理([参考](https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-by-leetcode/))：他的方法依赖于这样一个事实，在给定数组中最长上升子序列可能达到 `i^th`，独立于后面在数组中出现的元素。因此，如果我们知道 `lis` 的长度不超过 `i^th`，我们可以根据索引为`j` 的元素包括 `(i+1)^th`元素来计算 `lis` 的长度，其中 `0≤j≤（i+1）`

实现：新建数组`arr`，用于保存最长上升子序列，然后遍历`nums`，对于每个元素
- 如果 `arr` 中元素都比它小，将它插到最后
- 否则，用它覆盖掉比它大的元素中最小的那个，具体实现采用二分搜索
总之就是让`arr`中存储比较小的元素。这样，`arr`未必是真实的最长上升子序列，但长度是对的

代码:
``` golang
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	arr := make([]int, 0, len(nums))
	arr = append(arr, nums[0])
	for _, v := range nums[1:] {
		if v > arr[len(arr)-1] {
			arr = append(arr, v)
		} else {
			l, r := 0, len(arr)-1
			for l < r {
				mid := (l + r) / 2
				if arr[mid] >= v {
					r = mid
				} else {
					l = mid + 1
				}
			}
			arr[l] = v
		}
	}
	return len(arr)
}
```