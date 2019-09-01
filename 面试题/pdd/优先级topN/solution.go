package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	str := string(bytes)
	nums := []int{}
	start := 0
	for i := range str {
		if str[i] == ',' || str[i] == ';' {
			num, _ := strconv.Atoi(str[start:i])
			nums = append(nums, num)
			start = i + 1
		}

	}
	n, _ := strconv.Atoi(str[start:])
	heap := nums[:n]
	for i := n/2 - 1; i >= 0; i-- {
		shiftDown(heap, i, n)
	}
	l := len(nums)
	for i := n; i < l; i++ {
		tmp := max(nums[i], heap[0])
		if tmp == nums[i] {
			heap[0] = nums[i]
			shiftDown(heap, 0, n)
		}
	}
	for i := n - 1; i >= 0; i-- {
		heap[0], heap[i] = heap[i], heap[0]
		shiftDown(heap, 0, i)
	}
	for i := 0; i < n-1; i++ {
		fmt.Print(heap[i], ",")
	}
	fmt.Println(heap[n-1])
}

func shiftDown(arr []int, i int, n int) {
	left := (i * 2) + 1
	for left < n {
		right := left + 1
		tmp := left
		if right < n {
			tmp = max(arr[right], arr[left])
			if tmp == arr[right] {
				tmp = left
			} else {
				tmp = right
			}
		}
		temp := max(arr[tmp], arr[i])
		if temp == arr[i] {
			arr[tmp], arr[i] = arr[i], arr[tmp]
			i = tmp
			left = (i * 2) + 1
		} else {
			break
		}
	}
}

func max(i, j int) int {
	ii, jj := i%2, j%2
	if ii == 0 {
		if jj == 0 {
			if j > i {
				return j
			}
			return i
		}
		return i
	}
	if jj == 0 {
		return j
	}
	if j > i {
		return j
	}
	return i
}
