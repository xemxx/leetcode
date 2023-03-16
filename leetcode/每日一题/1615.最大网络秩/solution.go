// @algorithm @lc id=1738 lang=golang
// @title maximal-network-rank

package main

import "sort"

// @test(4,[[0,1],[0,3],[1,2],[1,3]])=4
// @test(5,[[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]])=5
// @test(8,[[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]])=5
// @test(2,[[1,0]])=1
func maximalNetworkRank(n int, roads [][]int) int {
	// 纯暴力
	t := make([]*M, n+1)
	for i := 0; i <= n; i++ {
		t[i] = &M{
			n: i,
			m: make(map[int]struct{}),
		}
	}
	for _, v := range roads {
		t[v[0]].m[v[1]] = struct{}{}
		t[v[1]].m[v[0]] = struct{}{}
	}
	sort.Slice(t, func(i, j int) bool {
		if t[i] == nil {
			return true
		}
		if t[i] == nil {
			return false
		}
		return len(t[i].m) < len(t[j].m)
	})
	// 如果最大只有一个，就直接和第二梯队做比较，如果有多个就多个里面两两暴力匹配
	tmp := len(t[n-1].m) + len(t[n].m)
	if len(t[n].m) == len(t[n-1].m) {
		for i := n; i >= 0; i-- {
			if len(t[i-1].m) < len(t[i].m) {
				break
			}
			for j := i - 1; j >= 0; j-- {
				if _, ok := t[j].m[t[i].n]; !ok {
					return tmp
				}
				if len(t[j-1].m) < len(t[j].m) {
					break
				}
			}
		}
		return tmp - 1
	}
	for i := n - 1; i >= 0; i-- {
		tmp := len(t[i].m) + len(t[n].m)
		if _, ok := t[n].m[t[i].n]; !ok {
			return tmp
		}
		if len(t[i-1].m) < len(t[i].m) {
			break
		}
	}
	return tmp - 1
}

type M struct {
	n int
	m map[int]struct{}
}
