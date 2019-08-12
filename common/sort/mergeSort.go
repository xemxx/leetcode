package sort

// MergeSort .
func MergeSort(arr []int, low, hight int) []int {
	if low < hight {
		mid := (hight-low)/2 + low
		MergeSort(arr, low, mid)
		MergeSort(arr, mid+1, hight)
		merge(arr, low, mid, hight)
		return arr
	}
	return []int{}
}

func merge(arr []int, low, mid, hight int) {
	a := make([]int, hight-low)
	i, j := low, mid
	for i < mid && j < hight {
		if arr[i] > arr[j] {
			a = append(a, arr[j])
			j++
		} else {
			a = append(a, arr[i])
			i++
		}
	}
	for i < mid {
		a = append(a, arr[i])
		i++
	}
	for j < hight {
		a = append(a, arr[j])
		j++
	}
}

// MergeSort1 is another solution .
// 该方法新建多个底层数组，空间浪费 .
func MergeSort1(arr []int) []int {
	l := len(arr)
	if l > 2 {
		mid := l / 2
		la := MergeSort1(arr[:mid])
		ra := MergeSort1(arr[mid:])
		a := merge1(la, ra)
		return a
	} else if l == 2 {
		a := arr
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	} else {
		return []int{}
	}
}

func merge1(la, ra []int) []int {

	a := []int{}
	i, j := 0, 0
	for i < len(la) && j < len(ra) {
		if la[i] > ra[j] {
			a = append(a, ra[j])
			j++
		} else {
			a = append(a, la[i])
			i++
		}
	}
	for i < len(la) {
		a = append(a, la[i])
		i++
	}
	for j < len(ra) {
		a = append(a, ra[j])
		j++
	}
	return a
}
