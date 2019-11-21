package main

import "fmt"

import "math"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println(-1)
		return
	}

	f := make([]uint64, n+1)
	if n < 2 {
		fmt.Println(1)
		return
	}
	f[1], f[2] = 1, 1
	for i := 3; i <= n; i++ {
		if (math.MaxUint64 - f[i-1]) < f[i-2] {
			fmt.Println(-1)
			return
		}
		f[i] = f[i-1] + f[i-2]
	}

	fmt.Println(f[n])

}
