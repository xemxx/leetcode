package main

import "fmt"

func main() {
	var l int
	fmt.Scan(&l)
	arr:=make([]int,l)
	
	for i:=0;i<l;i++{
		fmt.Scan(&arr[i])
	}
	var sum int
	fmt.Scan(&sum)
	m:=map[int]int{}
	for i:=0;i<l;i++{
		tmp:=sum-arr[i]
		if _,ok:=m[tmp];ok{
			fmt.Println(m[tmp],i)
			return
		}
		m[arr[i]]=i
	}
	fmt.Println(-1,-1)
}