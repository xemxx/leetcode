package main

import "fmt"

func main() {
	var n, k, l, r int
	fmt.Scan(&n, &k, &l, &r)
	count := 0
	for i := l; i <= r; i++ {
		if i%k == 0 {
			count++
		}
	}
	
}
