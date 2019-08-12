package sort

// QuickSort .
func QuickSort(arr []int, left, right int) []int {
	if left < right {
		pivot := paritition(arr, left, right)
		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
	return arr
}

//最基础的分区方法
func paritition(arr []int, left, right int) int {
	temp, i, j := arr[left], left, right
	for i < j {
		for arr[j] >= temp && j > i {
			j--
		}
		for arr[i] <= temp && j > i {
			i++
		}
		if i < j {
			swap(arr, i, j)
		}
	}
	swap(arr, j, left)
	return j
}

//对最基础的分区方法改进
func paritition1(arr []int, left, right int) int {
	temp, i, j := arr[left], left, right
	for i < j {
		for arr[j] >= temp && j > i {
			j--
		}
		arr[i] = arr[j]
		for arr[i] <= temp && j > i {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = temp
	return i
}

//比较巧妙的方法
func paritition2(arr []int, left, right int) int {
	temp, j := arr[left], left+1
	for i := j; i <= right; i++ {
		if arr[i] < temp {
			swap(arr, i, j)
			j++
		}
	}
	swap(arr, left, j-1)
	return j - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
