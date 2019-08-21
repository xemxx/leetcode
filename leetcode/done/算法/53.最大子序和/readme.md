## 动态规划
遍历整个数组，维护一个`iMax`和一个`max`值，分别代表当前为止当前连续序列的最大值和所有连续序列最大值
至于怎么保证`iMax`的值是一个连续序列相乘得到的，只要使得：

``` golang
iMax = max(iMax*v, v)
```

这样得到的`iMax`要么是连续的序列相乘至此，要么是仅由一个值得来，即可保证是一个连续序列相乘得到。

然后每次需要更新一下所有连续序列的最大值`max`


完整代码：
``` golang

func maxSubArray(nums []int) int {
	max, iMax := nums[0], nums[0]
	for _, v := range nums[1:] {
		tmp := v + iMax
		if tmp < v {
			tmp = v
		}
		iMax = tmp
		if tmp > max {
			max = tmp
		}
	}
	return max
}
```