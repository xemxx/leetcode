package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	f := []int{0, 1, 1, 1, 1}
	for i := 5; i <= n; i++ {
		f = append(f, f[i-1]+f[i-4])
	}
	fmt.Println(f[n])
}
