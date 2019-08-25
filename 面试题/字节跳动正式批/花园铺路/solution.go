package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	arr := map[int]int{}
	arr[2] = 1
	arr[4] = 2
	for i := 6; i <= n; i += 2 {
		arr[i] = arr[i-2] * 2
		for j := 4; j < i; j += 2 {
			arr[i] += (arr[j-2] * arr[i-j]) % 1000000007
			arr[i] %= 1000000007
		}
	}
	fmt.Println(arr[n])
}
