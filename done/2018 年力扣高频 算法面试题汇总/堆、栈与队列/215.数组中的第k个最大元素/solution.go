package main

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

func findKthLargest(nums []int, k int) int {
	k--
	quickSort(nums, 0, len(nums)-1, k)
	return nums[k]
}

func quickSort(arr []int, left, right, k int) {
	if left < right {
		pivot := paritition(arr, left, right)
		if pivot > k {
			quickSort(arr, left, pivot-1, k)
		} else if pivot < k {
			quickSort(arr, pivot+1, right, k)
		}
	}
}

func paritition(arr []int, left, right int) int {
	temp, i, j := arr[left], left, right
	for i < j {
		for arr[j] <= temp && j > i {
			j--
		}
		for arr[i] >= temp && j > i {
			i++
		}
		if i < j {
			swap(arr, i, j)
		}
	}
	swap(arr, j, left)
	return i
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
