package main

import (
	"fmt"
	"sort"
)

type data [][]int

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, 4)
		fmt.Scan(&tmp[0], &tmp[1])
		tmp[2] = i
		arr[i] = tmp
	}
	can := len(arr)
	for can >= m {
		//吃
		tmp := append([][]int{}, arr...)
		sort.Sort(data(tmp))
		//fmt.Println(tmp)
		if arr[tmp[0][2]][0] != 0 {
			arr[tmp[0][2]][0] = -1
			can--
		} else {
			break
		}
		if arr[tmp[1][2]][0] != 0 {
			arr[tmp[1][2]][0] = -1
			can--
		} else {
			break
		}
		//fmt.Println(tmp)
		//扣
		for i := 0; i < n; i++ {
			if arr[i][0] > 0 {
				arr[i][0] -= arr[i][1]
				if arr[i][0] <= 0 {
					arr[i][0] = 0
					can--
				}
			}
		}
		//fmt.Println(arr)
	}

	for _, v := range arr {
		fmt.Println(v[0])
	}

}

func (d data) Len() int {
	return len(d)
}

func (d data) Less(i, j int) bool {
	if d[i][0] > d[j][0] {
		return true
	} else if d[i][0] == d[j][0] {
		if d[i][1] > d[j][0] {
			return false
		}
		return true
	} else {
		return false
	}
}

func (d data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
