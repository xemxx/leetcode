package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	fmt.Scanln(&n)
	val := []uint{}
	key := []string{}
	for {
		var v uint
		var k string
		_, err := fmt.Scanln(&k, &v)
		if err == io.EOF {
			break
		}

		tmp := -1
		for i, j := range key {
			if k == j {
				tmp = i
				break
			}
		}

		if tmp == -1 {
			if len(key) < n {
				key = append(key, k)
				val = append(val, v)
			} else {
				fmt.Println(key[0], val[0])
				copy(key[0:n-1], key[1:])
				copy(val[0:n-1], val[1:])
				key[len(key)-1] = k
				val[len(val)-1] = v

			}
		} else if val[tmp] < v {
			size := len(key)
			copy(key[tmp:size-1], key[tmp+1:size])
			copy(val[tmp:size-1], val[tmp+1:size])
			val[size-1] = v
			key[size-1] = k
		}
	}
}
