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

理解：比如数组：`1, 8, 2, 6, 4, 5`，我们依次遍历，用`dp[i]`作为辅助数组
- 首先是1，那么`dp={1}`
- 然后是8，那么`dp={1,8}`
- 然后是2，此时我们只需要`dp={1,2}`，即可，因为如果后面的数字大于8，那么肯定大于2，所以只需要知道当前最长子序列为`{1,2}`即可
- 然后是6，此时`dp={1,2,6}`
- 然后是4，此时`dp={1,2,4}`，与上面同理
- 然后是5，此时`dp={1,2,4,5}`
- 最后得到最长上升子序列`{1,2,4,5}`
- 实际上`dp[i]`表示长度为`i+1`的所有递增子序列中，结尾最小的数字（或者说代表的含义是所有长度为`i`的上升子序列的最后一位的最小值），只要知道了最小值，那么遇到更大的，插入即可，所以此处用到了二分查找，再举个例子说明为什么要用二分查找：
  - 例如数组`{1,8,2,6,5,9,4}`
  - 遍历到9时，都与上述相同，此时`dp={1,2,5,9}`，直到4，此时4应该代替5，因为它代表了`{1,2,5}`和`{1,2,4}`两个长度为3的子序列时第三个数字的最小值，但是长度为4的子序列`{1,2,5,9}`不受影响，因为整个数组的长度只会增加或者不变，所以就算改变了前面的数字，最长的长度也不会改变了。

实现：新建数组`arr`，用于保存数组中每一位，然后遍历`nums`，对于每个元素：
- 如果`arr[len(arr)-1]<nums[i]`，代表存在长度为`i+2`的上升子序列，直接插入即可
- 否则，采用二分搜索，找到插入的位置，代表当前为止，所有长度为`i+1`的上升子序列中，结尾最小的数字为`nums[i]`

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