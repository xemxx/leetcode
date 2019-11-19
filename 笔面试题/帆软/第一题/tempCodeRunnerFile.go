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