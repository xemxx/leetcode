## 题目
互联网技术日新月异，为了让大家能学习新知识，我们在办公区设置了图书角供同学们借阅。过了一段时间，我们发现还书的同学没有把书按照顺序摆放导致找书找书很麻烦。安排专人整理图书顺序太花时间，管理数据中心的“天巡”机器人号称可以帮忙，但它只会从最右侧开始把一摞书全部拿起来翻个个儿。

例如：

当前有5本书，书的顺序是1 2 3 5 4。机器人能按照指令把最后两本书拿起来，翻个个儿放回去，变成：1 2 3 4 5。

如果5本书的顺序是1 2 5 3 4，机器人能按指令把最右侧两本书3 4拿起来翻个个儿变成4 3放回去，整体顺序变为：1 2 5 4 3；再把右边三本书5 4 3拿起来翻个个儿放回去，整体顺序变成1 2 3 4 5。

你能否写出一个程序，算出“天巡”机器人对于n本书，最少要翻转多少次才能把顺序排好？

输入:
输入参数为数组pBooks（长度为n的整数数组，每个元素是一个正整型数字表示书的序号）
第一行表示书本数n；
接下来n行，每行一个整数，表示书的序号。
输出:
输出机器人最少要翻转书的次数

## 思路
这个思路我也不知道是否正确，我这样的解法只能保证能有一个结果数- -
思路如下：
- 首先分为两个部分，排好序的和没排好序的，虽然题目没说，但是我认为数的序号是连续的正整数
- 直接从头遍历数组，每次判断当前位置的序号是否归位，即`arr[i]==i`是否为`true`，如果没有归位，开始判断归位此值需要多少步数，这里只会存在两种情况：
  - `i`在未排序部分的末尾，那么只需要将其反转一次即可归位
  - `i`在未排序部分的中间，那么只需要将其反转两次即可归位，一次是反转到末尾，一次是反转归位
- 很明显每次步数可能为`0,1,2`，每次相加即可

反正题目给的俩个例子我都过了，，，