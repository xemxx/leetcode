## 题目
给定一个n行m列数组，并且指定其中三列分别为index,key,val，要求转换为一个新的行列，以index为主键，key列为新的列头，val为key对应的列值只和，其余列的数据按照第一条不重复主键的值，如果对应key列该主键在原数组中不存在对应的val，则为-1
最后将新行列中的所有数据相加得出最后值
例如：
输入：
```
5 4 0 2 3
1 2 1 1
1 1 2 2
2 1 2 2
2 3 3 5
1 3 2 5
```
得到新行列：
```
1 2 1 7 -1
2 1 -1 2 5
```
最后结果为：19

## 思路
我的思路有两种，但是全部超时，只能过66.7%，我估摸着我的思路时间复杂度也就最多`O(n*m)`，，，看看有没有大佬能教我一下。。。

但还是给出我的思路：

第一种：
- 首先录入所有数据，分别保存所有key，以及每一个主键对应的原数组中的index，key，val组成的结构体值
- 然后遍历allRow，首先求出特殊值的和，用一个map标记即可，然后再求所有对应key列的val和，也用一个map标记然后将长度相减*-1即可得出其余所有为-1的val相加

第一种代码：
``` go
    var n, m, index, key, val int
	fmt.Scan(&n, &m, &index, &key, &val)
	arr := make([][]int, n)
    allKey := map[int]int{}
    
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&tmp[j])
        }
        // 保存需要的值
		allKey[tmp[key]] = 1
		allRow[tmp[index]] = append(allRow[tmp[index]], row{i: i, key: tmp[key], val: tmp[val]})
		arr[i] = tmp
	}
    result := 0
	for i, v := range allRow {
		tmp := i
		//非特殊值
		for j := 0; j < m; j++ {
			if j != index && j != key && j != val {
				tmp += arr[v[0].i][j]
			}
		}
		//特殊值
		tt := map[int]bool{}
		for _, row := range v {
			tmp += row.val
			tt[row.key] = true
		}
		tmp -= (len(allKey) - len(tt))
		result += tmp
    }
    fmt.Println(result)
```

第二种：
- 同样的首先录入所有参数，但是这次只需要所有的key即可
- 然后遍历每一行数据，用一个map数组标记新数组中唯一主键的行，用一个二维map标记每一个index拥有的key列，然后有两种情况：
  - 第一种是已经存在，则只需要将对应key下的所有值相加，并且将对应的key添加即可
  - 第二种是不存在，将对应的map标记为存在，并且将对应的key添加，然后找出所有非key列的值相加（因为这代表是新数组的唯一数据）
- 最后得到标记好每一个index的key数组`map[index][key]bool`，遍历所有index，将对应不存在的key列设为-1并且加在最后的结果上，则可以得到最后的总和值

代码：
``` go 
//录入代码略，参考上一个
    result:=0
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
		result -= len(allKey) - len(v)
    }
    fmt.Println(result)
```

其实第一种和第二种都差不太多，，，其中第二种我考虑过先排序然后就不需要再遍历一次key数组了，但是实际上我觉得这种标记后需要的时间更少