## 暴力法

循环每一个节点然后深搜所有可能，保存最大长度即可

## 排序后遍历

用时间复杂度为`O(nlgn)`的排序算法，例如归并排序，堆排序等

然后遍历，当遍历`nums[i]`时，判断与`nums[i-1]`的关系，如果`nums[i] == nums[i-1]+1`，则当前长度+1，否则更新最大长度，再将当前长度置1

## 用哈希表优化暴力法
解法：

这个优化算法与暴力算法仅有两处不同：

- 这些数字用一个`map[int]bool` 保存（其他语言用set或者hashset），实现O（1）的时间查询
- 遍历`map`得到每一个`num`值，判断`num-1`是否存在于`map`中，如果不存在则开始以此值开始寻找以此值开始的最长自序列，寻找则用到上一条的`map`即可，然后更新最大长度。如果存在`num-1`则不需要遍历`num`因为`num`所在的连续序列已经遍历过了

代码：
``` golang
func longestConsecutive(nums []int) int {
	mmp := make(map[int]bool,len(nums))
	for _, k := range nums {
		mmp[k] = true
	}
	max := 0
	for k := range mmp {
		if _, err := mmp[k-1]; err {
			cur := 1
			for mmp[k+1] != false {
				cur++
				k++
			}
			if cur > max {
				max = cur
			}
		}
	}
	return max
}
```

## 动态规划

- 用哈希表存储每个端点值对应连续区间的长度
- 若数已在哈希表中：跳过不做处理
- 若是新数加入：
    - 取出其左右相邻数已有的连续区间长度 left 和 right
    - 计算当前数的区间长度为：cur_length = left + right + 1
    - 根据 cur_length 更新最大长度 max_length 的值
    - 更新区间两端点的长度值
