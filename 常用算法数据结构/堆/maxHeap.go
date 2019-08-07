package heap

type maxHeap struct {
	data []int
}

func (c *maxHeap) Size() int {
	return len(c.data)
}

func (c *maxHeap) _shiftUp() {
	i := len(c.data) - 1
	for c.data[i] > c.data[i/2] && i > 1 {
		c.data[i], c.data[i/2] = c.data[i/2], c.data[i]
		i = i / 2
	}
}
func (c *maxHeap) Push(x int) {
	if c.data == nil {
		c.data = []int{0}
	}
	c.data = append(c.data, x)
	c._shiftUp()
}

func (c *maxHeap) _shiftDown() {
	for k := 1; k*2 <= len(c.data); {
		j := k * 2
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

func (c *maxHeap) Pop() int {
	if len(c.data) <= 1 {
		return 0
	}
	temp := c.data[1]
	c.data[1], c.data[len(c.data)-1] = c.data[len(c.data)-1], 0
	c.data = c.data[:len(c.data)-1]
	c._shiftDown()
	return temp
}
