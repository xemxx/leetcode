package main

type Tree struct {
	v  int
	ln *Tree
	rn *Tree
}

var num = map[int]int{}

func main() {
	num = make(map[int]int, 10)
	dp(&Tree{}, 0)
}

func dp(node *Tree, i int) {
	num[i] = node.v
	if node.rn != nil {
		dp(node.rn, i+1)
	}
	if node.ln != nil {
		dp(node.ln, i+1)
	}
}
