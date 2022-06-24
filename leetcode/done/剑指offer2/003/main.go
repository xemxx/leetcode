package main

import "fmt"

func main() {
	a := []int{2, 5, 16}
	for i := range a {
		fmt.Println(countBits(a[i]))
	}
}

// 11001 将第1位去掉，求剩下数字转换的十进制数对应有多少个1再加上第一位的一个1，就得到了所有的结果
// 11001 = 1+ res[1001]
// 或者直接右移一位，得到剩下数字转换的十进制对应有多少个1再看最后一位到底是不是1来选择是不是加1
// 11001 = res[1100] + 11001 & 1
func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	t := 1
	for i := 2; i <= n; i++ {
		if i == t<<1 {
			res[i] = 1
			t = t << 1
		} else {
			res[i] = 1 + res[i-t]
		}
	}
	return res
}
