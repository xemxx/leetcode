package main

import "fmt"

import "math"

func main() {
	var b int32
	fmt.Scan(&b)
	var result int32
	result = 0
	if b > 0 {
		for b != 0 {
			tmp := b % 10
			if result*10 > math.MaxInt32-tmp {
				result = 0
				break
			}
			result = result*10 + tmp
			b /= 10
		}
	} else if b < 0 {
		for b != 0 {
			tmp := b % 10
			if result*10 < math.MinInt32-tmp {
				result = 0
				break
			}
			result = result*10 + tmp
			b /= 10
		}
	}

	fmt.Println(result)
}
