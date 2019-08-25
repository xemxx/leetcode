package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)
	m := [][]int{}
	for i := 0; i <= 3; i++ {
		arr := make([]int, 4)
		for j := 0; j <= 3; j++ {
			fmt.Scan(&arr[j])
		}
		m = append(m, arr)
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
		j := r
		for j = r; 0 <= j && j <= 3; j += w {
			if n == 3 || n == 4 {
				if m[i][j] == 0 {
					for y := j; 0 < y && y < 3; y += w {
						m[i][y] = m[i][y+w]
					}
					m[i][3] = 0
				} else if m[i][j] == m[i][j+1] {
					m[i][j+1] = m[i][j] + m[i][j+1]
					for y := j; 0 < y && y < 3; y += w {
						m[i][y] = m[i][y+w]
					}
					m[i][3] = 0
				}
			} else {
				if m[j][j] == 0 {
					for y := j; 0 < y && y < 3; y += w {
						m[y][i] = m[y+w][i]
					}
					m[3][i] = 0
				} else if m[j][i] == m[j+1][i] && j+1 != j {
					m[j+1][i] = m[j][i] + m[j+1][i]
					for y := j; 0 < y && y < 3; y += w {
						m[y][i] = m[y+w][i]
					}
					m[3][i] = 0
				}
			}
		}
		j = r
	}
	for i := range m {
		s := ""
		for _, j := range m[i] {
			s += strconv.Itoa(j) + " "
		}
		fmt.Println(s[:len(s)-1])
	}
}
