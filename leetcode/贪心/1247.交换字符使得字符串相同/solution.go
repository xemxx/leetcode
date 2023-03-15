// @algorithm @lc id=1369 lang=golang
// @title minimum-swaps-to-make-strings-equal

package main

// @test("xx","yy")=1
// @test("xy","yx")=2
// @test("xx","xy")=-1
func minimumSwap(s1 string, s2 string) int {
	cntX, cntY := 0, 0
	result := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				if cntX == 0 {
					cntX++
					result++
				} else {
					cntX = 0
				}
			} else {
				if cntY == 0 {
					cntY++
					result++
				} else {
					cntY = 0
				}
			}
		}
	}
	if cntX != cntY {
		return -1
	}
	return result
}
