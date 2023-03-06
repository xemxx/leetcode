// @algorithm @lc id=455 lang=golang
// @title assign-cookies

package main

import "sort"

// @test([1,2,3],[1,1])=1
// @test([1,2],[1,2,3])=2
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			count++
			i++
		}
		j++
	}
	return count
}
