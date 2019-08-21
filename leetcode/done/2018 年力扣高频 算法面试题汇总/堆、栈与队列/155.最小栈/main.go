package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 24, 5, 3, 56, 7, 9, 2}
	obj := Constructor()
	for _, v := range arr {
		obj.Push(v)
	}
	obj.Pop()
	a1 := obj.Top()
	a2 := obj.GetMin()
	fmt.Println(a1,a2)
}
