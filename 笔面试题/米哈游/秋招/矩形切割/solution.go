package main

import "fmt"

func main() {
	var h, l int
	fmt.Scan(&l, &h)
	var n int
	fmt.Scan(&n)
	x := []int{0}
	y := []int{0}
	for i := 0; i < n; i++ {
		var t string
		var v int
		fmt.Scan(&t, &v)
		if t == "H" {
			// y = append(y, v)
			// for j := len(y) - 1; j > 0; j-- {
			// 	if y[j] < y[j-1] {
			// 		y[j-1], y[j] = y[j], y[j-1]
			// 	}
			// }
			if len(y) == 0 || v > y[len(y)-1] {
				y = append(y, v)
			} else {
				l, r := 0, len(y)-1
				for l < r {
					mod := (r-l)/2 + l
					if y[mod] >= v {
						r = mod
					} else {
						l = mod + 1
					}
				}
				y = append(y, 0)
				copy(y[l+1:], y[l:])
				y[l] = v
			}
		} else {
			// x = append(x, v)
			// for j := len(x) - 1; j > 0; j-- {
			// 	if x[j] < x[j-1] {
			// 		x[j-1], x[j] = x[j], x[j-1]
			// 	}
			// }
			if len(x) == 0 || v > x[len(x)-1] {
				x = append(x, v)
			} else {
				l, r := 0, len(x)-1
				for l < r {
					mod := (r-l)/2 + l
					fmt.Println(r, l, mod, x)
					if x[mod] >= v {
						r = mod
					} else {
						l = mod + 1
					}

				}
				x = append(x, 0)
				//fmt.Println(r, l, x)
				copy(x[l+1:], x[l:])
				x[l] = v
				fmt.Println(r, l, x)
			}
		}
	}
	y = append(y, h)
	x = append(x, l)
	max := 0

	for i := 1; i < len(x); i++ {
		for j := 1; j < len(y); j++ {
			now := (x[i] - x[i-1]) * (y[j] - y[j-1])
			if now > max {
				max = now
			}
		}
	}
	fmt.Println(max)
}
