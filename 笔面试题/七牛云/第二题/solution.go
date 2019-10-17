package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	arr:=[]string{}
	for i:=0;i<len(s);i++{
		c:=s[i:i+1]
		if c =="(" || c=="{" || c == "[" || c=="<" {
			arr=append(arr,c)
		}else{
			if len(arr)>0{
				d:=arr[len(arr)-1]
				if (c==")" && d=="(") || (c=="}" && d=="{") || (c=="]" && d=="[") || (c==">" && d=="<") {
					arr=arr[:len(arr)-1]
				}else{
					fmt.Println(0)
					return 
				}
			}else{
				fmt.Println(0)
				return 
			}
		}
		//fmt.Println(c,arr)
	}
	fmt.Println(1)
}