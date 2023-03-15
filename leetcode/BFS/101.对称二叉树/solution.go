// @algorithm @lc id=101 lang=golang
// @title symmetric-tree

package main

// @test([1,2,2,3,4,4,3])=true
// @test([1,2,2,null,3,null,3])=false
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	result := []int{}

	cur := []*TreeNode{root}
	nxt := []*TreeNode{}

	for len(cur) > 0 {
		n := cur[0]
		cur = cur[1:]
		if n == nil {
			result = append(result, -101)
		} else {
			nxt = append(nxt, n.Left, n.Right)
			result = append(result, n.Val)
		}
		if len(cur) == 0 {
			//check
			if len(result) > 1 {
				for i, j := 0, len(result)-1; i < j; {
					if result[i] != result[j] {
						return false
					}
					i++
					j--
				}
			}
			result = result[:0]
			cur, nxt = nxt, cur
			nxt = nxt[:]
			nxt = nxt[:0]
			if len(cur)%2 > 0 {
				return false
			}
		}
	}
	return true
}
