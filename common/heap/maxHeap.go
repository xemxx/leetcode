package heap

// 主要是实现最大堆，最小堆直接将大小判断反过来即可，不与实现
// 此外堆排序位于sort包内，同时推荐阅读go的官方sort包

//MaxHeap .
type MaxHeap struct {
	data []int
}

// Insert .
func (c *MaxHeap) Insert(x int) {
	if c.data == nil {
		c.data = []int{}
	}
	c.data = append(c.data, x)
	c.shiftUp()
}

func (c *MaxHeap) shiftUp() {
	i := len(c.data) - 1
	for c.data[i] > c.data[i/2] && i > 1 {
		c.data[i], c.data[i/2] = c.data[i/2], c.data[i]
		i = i / 2
	}
}

//ExtractMax .
func (c *MaxHeap) ExtractMax() int {
	if len(c.data) < 1 {
		return 0
	}
	temp := c.data[0]
	c.data[0], c.data[len(c.data)-1] = c.data[len(c.data)-1], 0
	c.data = c.data[:len(c.data)-1]
	c.shiftDown()
	return temp
}

func (c *MaxHeap) shiftDown() {
	for k := 0; k*2+1 < len(c.data); {
		j := k*2 + 1
		if j+1 < len(c.data) {
			if c.data[j+1] > c.data[j] {
				j = j + 1
			}
		}
		if c.data[j] > c.data[k] {
			c.data[j], c.data[k] = c.data[k], c.data[j]
		}
		k = j
	}
}

// 上方是构建一个maxHeap结构并且实现基本的两个操作，Insert，ExtractMax
// 下面实现Create操作

// Create One MaxHeap  时间复杂度为O(N).
// 时间复杂度解释参考https://visualgo.net/zh/heap?slide=7-2 .
func Create(arr []int) []int {
	l := len(arr)
	for i := (l - 1) / 2; i >= 0; i-- {
		shiftDown(arr, i, l)
	}
	return arr
}

func shiftDown(arr []int, root, maxLen int) {
	for root*2+1 < maxLen {
		child := root*2 + 1
		if child+1 < maxLen {
			if arr[child+1] > arr[child] {
				child = child + 1
			}
		}
		if arr[child] <= arr[root] {
			break
		}
		arr[child], arr[root] = arr[root], arr[child]
		root = child
	}
}

// Create1 时间复杂度为O(NlogN).
func Create1(arr []int) MaxHeap {
	l := len(arr)
	res := MaxHeap{make([]int, l)}
	for v := range arr {
		res.Insert(v)
	}
	return res
	// 或者如下： 注意修改函数返回值
	// l := len(arr)
	// for i := 0; i < l; i++ {
	// 	shiftDown(arr, i, l)
	// }
	// return arr
}
