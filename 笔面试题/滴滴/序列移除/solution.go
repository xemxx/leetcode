package main

import "fmt"

func main() {
	var n, total, cost int
	fmt.Scanln(&n, &total, &cost)
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	count, costs := dp(a, b, total, cost)
	count2 := dp2(a, b, total, cost)
	fmt.Println(count, costs, count2)
}

func dp(a, b []int, total, cost int) (int, int) {
	if len(a) == 0 || len(b) == 0 {
		return 0, len(b) + len(a)
	}
	am, bm := map[int]int{}, map[int]int{}
	max := len(a)
	if len(b) > max {
		max = len(a)
	}
	count, costs := 0, 0
	for i := 0; i < max; i++ {
		am[a[i]], bm[b[i]] = i, i
		if v, ok := bm[a[i]]; ok {
			count, costs = dp(a[i+1:], b[v+1:], total, cost)
			break
		}
		if v, ok := am[b[i]]; ok {
			count, costs = dp(a[v+1:], b[i+1:], total, cost)
			break
		}
	}
	if costs+cost > total {
		return count, costs
	}
	return count + 1, costs + cost
}

func dp2(a, b []int, total, cost int) int {
	n := len(a)
	am, bm := map[int]int{}, map[int]int{}
	count := 0
	for i := 0; i < n; i++ {
		am[a[i]], bm[b[i]] = i, i
		last := 0
		if v, ok := bm[a[i]]; ok && v != -1 {
			for j := i; j >= 0; j-- {
				am[a[j]] = -1
			}
			for j := v; j >= 0; j-- {
				bm[b[j]] = -1
			}
			count++
			total -= cost
			last += (n - i) + (n - v)
		}
		if v, ok := am[b[i]]; ok && v != -1 {
			for j := v; j >= 0; j-- {
				am[a[j]] = -1
			}
			for j := i; j >= 0; j-- {
				bm[b[j]] = -1
			}
			count++
			total -= cost
			last += (n - i) + (n - v)
		}
		if total-last < cost {
			count--
			break
		}
	}
	return count
}
