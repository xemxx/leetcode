package main
import (
	"fmt"
)
var max int
func main() {
	var t int
	fmt.Scan(&t)
	for i:=0;i<t;i++{
		max=0
		var a,b int
		fmt.Scanln(&a,&b)
		if b%2!=0 && (b-1)%10!=0{
			fmt.Println(-1)
		}else{
			fmt.Println(dp(a,b,0))
		}
	}
}

func dp(a,b,now int) int {
	if a>b{
		return -1
	}else if a==b{
		return now
	}
	m:=dp(a*2,b,now+1)
	n:=dp(a*10+1,b,now+1)
	if m==-1{
		return n
	}else if n==-1{
		return m
	}else if n<m{
		n=m
	}
	return m
}

func get(a,b,now int) int{
	if a>b{
		return -1
	}else if a==b{
		return now
	}
	m:=dp(a*10+1,b,now+1)
	if m==-1{
		n:=dp(a*2,b,now+1)
		if n==-1{
			return -1
		}
		return n
	}
	return m
}