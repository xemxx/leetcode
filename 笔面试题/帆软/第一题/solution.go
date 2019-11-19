package main

import "fmt"

type row struct {
	i   int
	key int
	val int
}

func main() {
	var n, m, index, key, val int
	result := 0
	fmt.Scan(&n, &m, &index, &key, &val)
	arr := make([][]int, n)
	allKey := map[int]int{}
	//allRow := map[int][]row{}
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		//tt := 0
		for j := 0; j < m; j++ {
			fmt.Scan(&tmp[j])
			// if j != index && j != key && j != val {
			// 	tt += tmp[j]
			// }
		}
		allKey[tmp[key]] = 1
		// if _, ok := allRow[tmp[index]]; !ok {
		// 	result += tt
		// }
		// if _, ok := allRow[tmp[index]]; !ok {
		// 	for j := 0; j < m; j++ {
		// 		if j != index && j != key && j != val {
		// 			result += tmp[j]
		// 		}
		// 	}
		// }
		//allRow[tmp[index]] = append(allRow[tmp[index]], row{i: i, key: tmp[key], val: tmp[val]})
		arr[i] = tmp
	}

	// result := [][]int{}

	// for i, v := range allRow {
	// 	tmp := i
	// 	//非特殊值
	// 	for j := 0; j < m; j++ {
	// 		if j != index && j != key && j != val {
	// 			tmp += arr[v[0].i][j]
	// 		}
	// 	}
	// 	//fmt.Println(i, v, tmp)
	// 	//特殊值
	// 	tt := map[int]bool{}
	// 	for _, row := range v {
	// 		tmp += row.val
	// 		tt[row.key] = true
	// 	}
	// 	tmp -= (len(allKey) - len(tt))
	// 	result += tmp
	// }

	rows := map[int]bool{}
	tt := map[int](map[int]bool){}
	for i := 0; i < n; i++ {
		if _, ok := rows[arr[i][index]]; ok {
			tt[arr[i][index]][arr[i][key]] = true
			result += arr[i][val]
		} else {
			rows[arr[i][index]] = true
			tt[arr[i][index]] = map[int]bool{}
			tt[arr[i][index]][arr[i][key]] = true
			for j := 0; j < m; j++ {
				if j != key {
					result += arr[i][j]
				}
			}
		}
	}
	for _, v := range tt {
		//fmt.Println(result)
		result -= len(allKey) - len(v)
	}

	// for j := 0; j < m; j++ {
	// 	if j != index && j != key && j != val {
	// 		tmp := 0
	// 		rows := map[int]bool{}
	// 		for i := 0; i < n; i++ {
	// 			if _, ok := rows[arr[i][index]]; !ok {
	// 				rows[arr[i][index]] = true
	// 				tmp += arr[i][j]
	// 			}
	// 		}
	// 		result += tmp
	// 	}
	// }
	fmt.Println(result)

}
