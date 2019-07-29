package main

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack []*TreeNode

func postorderTraversal(root *TreeNode) []int {
	re := []int{}
	arr := stack{}
	node, lv := root, &TreeNode{}
	for node != nil || len(arr) > 0 {
		if node != nil {
			arr.push(node)
			node = node.Left
		} else {
			temp := arr.pop()
			if temp.Right == lv || temp.Right == nil {
				re = append(re, temp.Val)
				node = nil
				lv = temp
			} else {
				arr.push(temp)
				node = temp.Right
			}
		}
	}
	return re
}
func (r *stack) push(x *TreeNode) {
	(*r) = append((*r), x)
}

func (r *stack) pop() *TreeNode {
	x := (*r)[len(*r)-1]
	(*r)[len(*r)-1] = nil
	*r = (*r)[:len(*r)-1]
	return x
}
