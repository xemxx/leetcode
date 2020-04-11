package main

import (
	"fmt"
)

func main() {
	fmt.Println("finish")
}

// 1.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsTreeMeetTarget(root *TreeNode, target int) bool {
	if root == nil || root.Val > target {
		return false
	} else if root.Val == target {
		return true
	}
	return IsTreeMeetTarget(root.Left, target-root.Val) || IsTreeMeetTarget(root.Right, target-root.Val)
}

// 2.假定是升序
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoSortedList(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tmp.Next = &ListNode{l2.Val, nil}
			tmp = tmp.Next
			l2 = l2.Next
		} else {
			tmp.Next = &ListNode{l1.Val, nil}
			tmp = tmp.Next
			l1 = l1.Next
		}
	}
	if l1 != nil {
		for l1 != nil {
			tmp.Next = &ListNode{l1.Val, nil}
			tmp = tmp.Next
			l1 = l1.Next
		}
	} else if l2 != nil {
		for l2 != nil {
			tmp.Next = &ListNode{l2.Val, nil}
			tmp = tmp.Next
			l2 = l2.Next
		}
	}
	return result.Next
}

// 3.
func threeSum(nums []int) [][]int {
	result := [][]int{}
	hm := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hm[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		tmp := make([]int, 3)
		t1 := 0 - nums[i]
		yes := false
		for j := i + 1; j < len(nums); j++ {
			t2 := t1 - nums[j]
			if v, ok := hm[t2]; ok && v > j {
				tmp[2] = t2
				tmp[1] = nums[j]
				yes = true
				break
			}
		}
		if yes {
			tmp[0] = nums[i]
			result = append(result, tmp)
		}
	}
	return result
}

// 4.
func partitionLabels(s string) []string {
	allMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := allMap[s[i]]; ok {
			allMap[s[i]]++
		}
		allMap[s[i]] = 1
	}
	tmp := map[byte]int{}
	result := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if start == i {
			tmp[s[i]] = 1
		} else {
			if _, ok := tmp[s[i]]; ok {
				tmp[s[i]]++
			}
			tmp[s[i]] = 1
		}
		yes := false
		for k, v := range tmp {
			if allMap[k] == v {
				yes = true
			} else {
				yes = false
			}
		}
		if yes {
			result = append(result, s[start:i+1])
			start = i + 1
			tmp = map[byte]int{}
		}
	}
	return result
}
