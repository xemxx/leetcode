package main

import "fmt"

func main() {
	fmt.Println(solve([]string{"aa", "bb", "cc", "dd", "eeee"}, 2))
}

func solve(arr []string, flag int) [][]string {
	result := [][]string{}
	j := 1
	tmp := []string{}
	for i := 0; i < len(arr); i++ {
		if j < flag {
			tmp = append(tmp, arr[i])
			j++
		} else if j == flag {
			tmp = append(tmp, arr[i])
			result = append(result, tmp)
			j = 1
			tmp = []string{}
		}
	}
	if j != 1 {
		result = append(result, tmp)
	}
	return result
}

// 用户，角色，每个角色权限，权限对应功能
