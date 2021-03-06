主要采用动态规划
遍历整个数组，维护一个iMax和一个iMin值，分别代表当前为止最大值和最小值
至于怎么保证iMax或者iMin的值是一个连续序列相乘得到的，只要使得：

``` golang
iMax = max(iMax*v, v)
//同理iMin
iMin = min(iMin*v, v)
```

这样得到的iMax要么是连续的序列相乘至此，要么是仅由一个值得来，即可保证是一个连续序列相乘得到。

同时因为涉及到乘法的正负问题，遇到负数时，最小的值将变成最大，最大值将变成最小，所以就需要同时保存最小值和最大值，同时相应的处理如下：

``` golang
if v < 0 {
    iMax, iMin = iMin, iMax
}
```
