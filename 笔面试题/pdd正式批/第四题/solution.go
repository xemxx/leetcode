package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i:=0;i<t;i++{
		var n,m int
		fmt.Scanln(&n,&m)
		a:=make([]int,n)
		for j:=0;j<n;j++{
			fmt.Scan(&a[j])
		}
		solve(a,n,m)
	}
}

func solve(arr []int,n,m int){
	mm:=map[int]bool{}
	count:=1
	mm[arr[0]]=true
	for i:=1;i<n;i++{
		if arr[i]==arr[i-1]{
			if m>1{
				mm=map[int]bool{}
			}
		}
		if _,ok:=mm[arr[i]];!ok{
			mm[arr[i]]=true
			count++
		}
	}
	fmt.Println(count)
	
}