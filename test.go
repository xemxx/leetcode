package main

import "fmt"

func main() {
	s := "1245152451321211"
	dp(s, 6, "")
}
func dp(s string, n int, re string) {
	ls := len(s)
	if ls<=n*3 && ls>n{
		if n == 1 {
			fmt.Println(re + s)
		}
		for i:=1;i<=ls && i<=3;i++{
			dp(s[i:],n-1,re+s[:i]+"*") 
		}
	}
}
