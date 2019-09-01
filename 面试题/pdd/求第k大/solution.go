package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scanln(&n, &m, &k)
	num := m - n + 1
	i := n
	if m < n {
		num = n - m + 1
		i = m
	}
	a := 0
	for k > (num + a) {
		k -= (num + a)
		num++
		a++
		i--
	}
	if k <= 0 {
		a--
		num--
		k += num + a
		i++
	}
	j := 0
	for k > 0 {
		k -= 2
		j++
	}
	res := 0
	j--
	if n <= m {
		res = i * (m - j)
	} else {
		res = i * (n - j)
	}
	fmt.Println(res)

}
