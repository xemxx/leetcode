package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(kthSmallest1(&TreeNode{3, &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}, 2))
}

//方法一 用中序遍历的方法生成有序数组 由于是递归所以时间复杂度较高
func kthSmallest(root *TreeNode, k int) int {
	arr := dfs(root)
	return arr[k-1]
}
func dfs(node *TreeNode) []int {
	arr := []int{}
	if node != nil {
		lNode := node.Left
		add := dfs(lNode)
		for _, v := range add {
			arr = append(arr, v)
		}
		arr = append(arr, node.Val)
		rNode := node.Right
		add = dfs(rNode)
		for _, v := range add {
			arr = append(arr, v)
		}

	}
	return arr
}

//方法二 用数组保存各个节点，然后按照中序遍历的顺序查找第k大的节点 时间复杂度更低
func kthSmallest1(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	ans := 0
	for true {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//fmt.Println(stack)
		if len(stack) == 0 {
			break
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//fmt.Println(node)
		ans = node.Val
		//fmt.Println(ans)
		k--
		if k == 0 {
			break
		}
		root = node.Right

	}
	return ans
}

//第三种 也是按照中序遍历的方式查找最少元素，为上一种方法的递归实现 但是在leetcode上花费时间更少
func kthSmallest2(root *TreeNode, k int) int {
	return *(traversal(root, &k))
}
func traversal(node *TreeNode, k *int) *int {
	if node.Left != nil {
		res := traversal(node.Left, k)
		if res != nil {
			return res
		}
	}
	(*k)--
	if (*k) == 0 {
		return &node.Val
	}
	if node.Right != nil {
		return traversal(node.Right, k)
	}
	return nil
}

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (63.18%)
 * Total Accepted:    10.7K
 * Total Submissions: 17K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
 *
 * 说明：
 * 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
 *
 * 示例 1:
 *
 * 输入: root = [3,1,4,null,2], k = 1
 * ⁠  3
 * ⁠ / \
 * ⁠1   4
 * ⁠ \
 * 2
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: root = [5,3,6,2,4,null,null,1], k = 3
 * ⁠      5
 * ⁠     / \
 * ⁠    3   6
 * ⁠   / \
 * ⁠  2   4
 * ⁠ /
 * ⁠1
 * 输出: 3
 *
 * 进阶：
 * 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
