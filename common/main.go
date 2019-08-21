package main

import (
	"fmt"
	"leetcode/common/heap"
	"leetcode/common/sort"
)

func main() {
	a := []int{3, 44, 38, 5, 47, 15, 36, 26, 26, 27, 2, 46, 4, 19, 50, 48, 49}
	fmt.Printf("QuickSort() = %v \n", sort.QuickSort(append([]int{}, a...), 0, len(a)-1))
	fmt.Printf("MergeSort() = %v \n", sort.MergeSort(append([]int{}, a...), 0, len(a)-1))
	fmt.Printf("MergeSort1() = %v \n", sort.MergeSort1(append([]int{}, a...)))
	fmt.Printf("HeapSort() = %v \n", sort.HeapSort(append([]int{}, a...)))
	fmt.Printf("MaxHeap.Create1() = %v \n", heap.Create1(append([]int{}, a...)))
	fmt.Printf("Orange Arr = %v \n", a)
}
