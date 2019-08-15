package io

import (
	"fmt"
	"io"
)

func scanArr() [][]int {
	j := 0
	var v []int
	arr := [][]int{[]int{}}
	for {
		_, err := fmt.Scanln(&v)
		if err == io.EOF {
			break
		}
		fmt.Println(v)
		for _, k := range v {
			arr[j] = append(arr[j], k)
		}
		arr = append(arr, []int{})
		j++
	}
	fmt.Println(arr)
	return arr
}
