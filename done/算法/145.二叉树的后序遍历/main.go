package main

import "fmt"

func main() {
	tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(postorderTraversal(tree))
}
