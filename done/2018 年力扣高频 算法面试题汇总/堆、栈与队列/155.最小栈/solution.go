package main

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

//通过维护两个数组，一个用于获取最小数
type MinStack struct {
	data  []int
	sData []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

/*
* push时，辅助栈有两种情况需要push
* 一是x<=当前辅助栈栈顶元素，二是为空时，
* 为啥x要判断等于的情况是因为可能有2个相同的元素，为了后续的pop操作
* 如果x>当前辅助栈栈顶元素自然就不需要入栈，因为如果x不为最小值且为data栈的栈顶元素时，一定存在更小的元素于data栈中，那么pop此元素和不pop此元素都是一样的结果
 */
func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.sData) == 0 || this.sData[len(this.sData)-1] >= x {
		this.sData = append(this.sData, x)
	}
}

func (this *MinStack) Pop() {
	x := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if x == this.sData[len(this.sData)-1] {
		this.sData = this.sData[:len(this.sData)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.sData) > 0 {
		return this.sData[len(this.sData)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
