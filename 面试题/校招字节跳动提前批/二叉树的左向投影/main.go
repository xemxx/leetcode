package main

type tree struct {
	v  int
	ln *tree
	rn *tree
}

var num = map[int]int{}

func main() {
	num = make(map[int]int, 10)
	dp(&tree{}, 0)
}

func dp(node *tree, i int) {
	num[i] = node.v
	if node.rn != nil {
		dp(node.rn, i+1)
	}
	if node.ln != nil {
		dp(node.ln, i+1)
	}
}
