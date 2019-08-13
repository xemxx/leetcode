package main

import "fmt"
import "leetcode/common/sort"

func main() {
	a := []int{3, 44, 38, 5, 47, 15, 36, 26, 26, 27, 2, 46, 4, 19, 50, 48, 49}
	fmt.Printf("QuickSort() = %v \n", sort.QuickSort(a, 0, len(a)-1))
	fmt.Printf("MergeSort() = %v \n", sort.MergeSort(a, 0, len(a)-1))
	fmt.Printf("HeapSort() = %v \n", sort.HeapSort(a))
}
