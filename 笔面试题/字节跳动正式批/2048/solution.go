package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)
	// m := [][]int{}
	// for i := 0; i <= 3; i++ {
	// 	arr := make([]int, 4)
	// 	for j := 0; j <= 3; j++ {
	// 		fmt.Scan(&arr[j])
	// 	}
	// 	m = append(m, arr)
	// }
	m := [][]int{
		{0, 0, 0, 2},
		{2, 1, 1, 2},
		{2, 0, 4, 8},
		{8, 0, 4, 8},
	}
	q, w := 1, 1
	e, r := 0, 0
	switch n {
	case 1:
		q, w, e, r = 1, 1, 0, 0
	case 2:
		q, w, e, r = 1, -1, 0, 3
	case 3:
		q, w, e, r = 1, 1, 0, 0
	case 4:
		q, w, e, r = 1, -1, 0, 3
	}
	for i := e; 0 <= i && i <= 3; i += q {
		x, j := r, r
		tail := 4
		pid := r + 3*w

		for j += w; tail > 0; tail-- {
			if n == 3 || n == 4 {
				if m[i][x] == 0 {
					y := 0
					for y = x; y != pid; y += w {
						m[i][y] = m[i][y+w]
					}
					m[i][y] = 0
				} else if m[i][j] == 0 {
					y := 0
					for y = j; y != pid; y += w {
						m[i][y] = m[i][y+w]
					}
					m[i][y] = 0
				} else if m[i][j] == m[i][x] {
					m[i][x] = m[i][j] + m[i][x]
					y := 0
					for y = j; 0 < y && y < 3; y += w {
						m[i][y] = m[i][y+w]
					}
					m[i][y] = 0
					x = j
					j += w
				} else {
					x = j
					j += w
				}
			} else {
				if m[x][i] == 0 {
					y := 0
					for y = x; y != pid; y += w {
						m[y][i] = m[y+w][i]
					}
					m[y][i] = 0
				} else if m[j][i] == 0 {
					y := 0
					for y = j; y != pid; y += w {
						m[y][i] = m[y+w][i]
					}
					m[y][i] = 0
				} else if m[j][i] == m[x][i] {
					m[x][i] = m[j][i] + m[x][i]
					y := 0
					for y = j; y != pid; y += w {
						m[y][i] = m[y+w][i]
					}
					m[y][i] = 0
					x = j
					j += w
				} else {
					x = j
					j += w
				}
			}
		}
	}
	for i := range m {
		s := ""
		for _, j := range m[i] {
			s += strconv.Itoa(j) + " "
		}
		fmt.Println(s[:len(s)-1])
	}
}
