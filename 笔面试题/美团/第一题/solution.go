package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := [][]int{}
	for j := 0; j < 2; j++ {
		var str string
		fmt.Scanf("%s\n", &str)
		// reader := bufio.NewReader(os.Stdin)
		// bytes, _, err := reader.ReadLine()
		// if err == io.EOF {
		// 	break
		// }
		// str := string(bytes)
		now := []int{}
		//fmt.Println(str)
		for _, v := range str {
			num := 0
			if v == 'X' {
				num = -1
			}
			now = append(now, num)
		}
		arr = append(arr, now)
	}
	//fmt.Println(arr)
	if arr[0][0] == -1 {
		fmt.Println("-1")
		return
	}
	arr[0][0] = 1
	for i := 1; i < n; i++ {
		count := 0
		for j := 0; j < 2; j++ {
			if arr[j][i] != -1 {
				tmp := 0
				if arr[0][i-1] != -1 {
					tmp += arr[0][i-1]
				}
				if arr[1][i-1] != -1 {
					tmp += arr[1][i-1]
				}
				arr[j][i] = tmp
			} else {
				count++
			}
			if count == 2 {
				fmt.Println("-1")
				return
			}
		}
	}
	if arr[1][n-1] > 0 {
		fmt.Println(arr[1][n-1])
	} else {
		fmt.Println("-1")
	}
}
