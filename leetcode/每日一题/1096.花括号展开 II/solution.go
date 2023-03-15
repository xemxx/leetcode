// @algorithm @lc id=1188 lang=golang
// @title brace-expansion-ii

package main

import (
	"sort"
	"strings"
)

// @test("{a,b}{c,{d,e}}")=["ac","ad","ae","bc","bd","be"]
// @test("{{a,z},a{b,c},{ab,z}}")=["a","ab","ac","z"]
func braceExpansionII(expression string) []string {
	m := map[string]struct{}{}
	dp(expression, m)
	result := []string{}
	for k := range m {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}

func dp(expression string, h map[string]struct{}) {
	i := strings.Index(expression, "}")
	if i == -1 {
		h[expression] = struct{}{}
		return
	}
	suffix := expression[i+1:]
	d := strings.LastIndex(expression[:i], "{")
	prefix := expression[:d]
	r := strings.Split(expression[d+1:i], ",")
	for _, v := range r {
		dp(prefix+v+suffix, h)
	}
}
