package main

import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	arr:=make([]int,n)
	
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	var num int
	fmt.Scan(&num)
	start:=-1
	end:=-1
	for i:=0;i<n;i++{
		if arr[i]==num{
			if i<start || start==-1{
				start=i
			}
			if i>end || end==-1{
				end=i
			}
		}else{
			if end>0{
				break
			}
		}
	}
	fmt.Println(start,end)
}