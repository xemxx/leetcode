package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(lowestCommonAncestor(&TreeNode{3, &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}))
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return &TreeNode{}
}

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
