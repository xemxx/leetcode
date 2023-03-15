package main

import "fmt"

type lru struct {
	Head *node
	Tail *node
	Map  map[int]int
	Num  int
	Now  int
}
type node struct {
	Key  int
	Val  int
	Next *node
	Prev *node
}

func main() {
	// l := NewLru(5)
	// l.add(1, 2)
	// l.add(2, 2)
	// l.add(3, 2)
	// l.pop(1)
	fmt.Println("94" < "9", "34" < "0")
}

func NewLru(b int) *lru {
	return &lru{nil, nil, map[int]int{}, b, 0}
}

func (l *lru) add(k, v int) {
	if l.Now < l.Num {
		l.Map[k] = v
		l.Tail.Next = &node{k, v, nil, l.Tail}
		l.Tail = l.Tail.Next
		l.Now++
	} else {
		l.Map[k] = v
		delete(l.Map, l.Head.Key)
		l.Tail.Next = &node{k, v, nil, l.Tail}
		l.Tail = l.Tail.Next
	}
}

func (l *lru) pop(k int) int {
	node := l.Head
	for node.Key != k {
		node = node.Next
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	l.Tail.Prev = node
	l.Tail = node
	return l.Map[k]
}
