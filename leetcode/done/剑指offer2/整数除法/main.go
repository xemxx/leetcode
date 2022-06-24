package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(15, 2))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(0, 1))
	fmt.Println(divide(1, 1))
	fmt.Println(divide(-2147483648, -1))
	fmt.Println(divide(-2147483648, -2147483648))
}
func divide(a int, b int) int {
	if a == 0 {
		return 0
	}
	if a == b {
		return 1
	}
	if b == math.MinInt32 {
		return 0
	}

	tmp := 1

	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		tmp = -1
	}
	if b > 0 {
		b = -b
	}

	if a > 0 {
		a = -a
	}

	result := 0
	if b == -1 {
		if tmp == 1 {
			result = -a

		} else {
			result = a
		}
	} else {
		now := b
		resulta, resultb := 0, tmp
		// 从分母入手，去掉已经除掉的部分，剩下的反复即可
		for a < 0 && a <= b {
			if now<<1 > a {
				resulta++
				now = now << 1
			} else {
				a = a - now
				now = b
				resultb = tmp
				for i := 0; i < resulta; i++ {
					resultb += resultb
				}
				resulta = 0
				result += resultb
			}
		}
	}
	if result > math.MaxInt32 {
		result = math.MaxInt32
	}
	return result
}
