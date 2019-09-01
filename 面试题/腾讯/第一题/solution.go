package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
	k := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&k[i])
	}

	k1, k2, t1, t2 := 0, 0, 0, 0
	for _, v := range t {
		if v%2 == 0 {
			t1++
		} else {
			t2++
		}
	}
	for _, v := range k {
		if v%2 == 0 {
			k1++
		} else {
			k2++
		}
	}
	if t1 > k2 {
		t1 = k2
	}
	if t2 > k1 {
		t2 = k1
	}
	fmt.Println(t1 + t2)
}
