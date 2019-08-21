package main

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

/*
* 每次遍历当前节点
* 如果当前节点不为nil，将当前节点push到stack中，并且切换当前节点为左子节点
* 如果当前节点为nil，即需要将当前节点的父节点值输出，并且遍历当前节点的父节点的右子节点，所以需要将之前保存过的父节点出栈，并且切换当前节点为父节点的右子节点
* stack的作用就是当左节点遍历完后需要返回到右节点继续遍历
 */

func inorderTraversal(root *TreeNode) []int {
	re := []int{}
	arr := stack{}
	node := root
	for node != nil || len(arr) > 0 {
		if node != nil {
			arr.push(node)
			node = node.Left
		} else {
			lnode := arr.pop()
			re = append(re, lnode.Val)
			node = lnode.Right
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
