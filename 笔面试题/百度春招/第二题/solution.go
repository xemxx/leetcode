package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(solve(n, k))
}

func solve(n, k int) int {
	if n <= k {
		return 1
	}
	tmp := (n - k) / 2
	if tmp*2 == (n-k) && tmp != 0 {
		return solve(tmp+k, k) + solve(tmp, k)
	}
	return 1
}
