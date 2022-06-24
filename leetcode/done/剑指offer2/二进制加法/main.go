package main

import "fmt"

func main() {
	a := []string{"1"}
	b := []string{"111"}
	for i := range a {
		fmt.Println(addBinary(a[i], b[i]))
	}
}

// 事实上二进制最多多进位一位数，所以这里只需要补齐前面的0，然后每次后面只会要么进1要么不进，不可能进更多，也就没有必要倒转了
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	res := make([]byte, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}

	res = append(res, '0')
	for i, j := len(b)-1, 0; i >= 0; i-- {
		tmp := b[i] - '0' + res[j]
		for tmp >= '2' {
			res[j+1] += 1
			tmp -= 2
		}
		res[j] = tmp
		j++
	}
	for i := len(b); res[i] >= '2'; i++ {
		tmp := byte('0')
		for res[i] >= '2' {
			tmp += 1
			res[i] -= 2
		}
		if len(res) > i+1 {
			res[i+1] = tmp - '0' + res[i+1]
		} else {
			res = append(res, tmp)
		}
	}
	if res[len(res)-1] == '0' {
		res = res[:len(res)-1]
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
