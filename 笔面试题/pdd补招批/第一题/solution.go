package main

import (
	"fmt"
	"sort"
)

type max [][]int
type min [][]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, 2)
		fmt.Scan(&tmp[0], &tmp[1])
		arr[i] = tmp
	}
	// maxi := max(arr)
	// mini := min(arr)
	sort.Sort(min(arr))

	trueMin, trueMax := -1, -1
	vmin, vmax := -1, -1
	for i := k - 1; i < n-k; i++ {
		tmp := 1
		for j := i - 1; j >= 0; j-- {
			if arr[i][0] <= arr[j][1] {
				tmp++
			}
		}
		//fmt.Println(mini[i], i, mini)
		if tmp == k {
			trueMin = i
			vmin = arr[i][0]
			break
		}
	}

	sort.Sort(max(arr))

	for i := k - 1; i < n-k; i++ {
		tmp := 1
		for j := i - 1; j >= 0; j-- {
			if arr[i][1] >= arr[j][0] {
				tmp++
			}
		}
		if tmp == k {
			trueMax = i
			vmax = arr[i][1]
			break
		}
	}
	if trueMin != -1 && trueMax != -1 {
		fmt.Printf("%v %v\n", vmin, vmax)
	} else {
		fmt.Println("error")
	}
}

func (d max) Len() int {
	return len(d)
}
func (d max) Less(i, j int) bool {
	if d[i][1] > d[j][1] {
		return true
	}
	return false
}

func (d max) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d min) Len() int {
	return len(d)
}
func (d min) Less(i, j int) bool {
	if d[i][0] < d[j][0] {
		return true
	}
	return false
}

func (d min) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
