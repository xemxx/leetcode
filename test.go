package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	a = a[1:3]
	fmt.Println(a)
}
func dp(s string, n int, re string) {
	ls := len(s)
	if ls <= n*3 && ls > n {
		if n == 1 {
			fmt.Println(re + s)
		}
		for i := 1; i <= ls && i <= 3; i++ {
			dp(s[i:], n-1, re+s[:i]+"*")
		}
	}
}
