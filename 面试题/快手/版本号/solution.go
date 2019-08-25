package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)
	a, b := make([]string, n), make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a[i], &b[i])
	}
	for i := 0; i < n; i++ {
		x := strings.Split(a[i], ".")
		y := strings.Split(b[i], ".")
		j, l := 0, len(x)
		if len(x) > len(y) {
			l = len(y)
		}
		res, is := false, 0
		for j = 0; j < l; j++ {
			n1, _ := strconv.Atoi(x[j])
			n2, _ := strconv.Atoi(y[j])
			if n1 < n2 {
				res = true
				is = 1
				break
			} else if n1 > n2 {
				res = false
				is = 1
				break
			} else {
				is = 2
			}
		}
		if is == 2 {
			nn := false
			if len(x) < len(y) {
				for j := range y {
					n1, _ := strconv.Atoi(y[j])
					if n1 > 0 {
						nn = true
						break
					}
				}

			}
			fmt.Println(nn)
		} else {
			fmt.Println(res)
		}

	}
}
