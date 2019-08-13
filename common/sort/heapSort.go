package sort

// HeapSort .
func HeapSort(arr []int) []int {
	l := len(arr)
	// 构建最大堆，采用时间复杂度为O(N)的方式构建
	// 时间复杂度解释参考https://visualgo.net/zh/heap?slide=7-2 .
	for i := (l - 1) / 2; i >= 0; i-- {
		shiftDown(arr, i, l)
	}
	// 堆排序具体实现，每次将堆顶元素与最末元素交换，并且堆长减一再shiftDown，最终实现升序排列
	for i := l - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown(arr, 0, i)
	}
	return arr
}

func shiftDown(arr []int, root, maxLen int) {
	// 此处不取等于的原因是因为maxLen是不可能取到的
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
