package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	//	zero, one, two := 1, 1, 1
	zp, zq := 1, 1
	op, oq := 0, 0
	tp, tq := 0, 0
	q := 1
	for j := 0; j < n; j++ {
		q *= a[j]
	}
	zq = q
	oq = q
	tq = q
	for i := 0; i < n; i++ {
		op += (a[i] - 1)
		tp += (a[i] - 1) * (i + 1)
		// for j := i + 1; j < n; j++ {
		// 	tp += (a[j] - 1)
		// }
	}
	fmt.Println((998244353+zp)/zq, (998244353+op)/oq, (998244353+tp)/tq)

}
