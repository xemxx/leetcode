// @algorithm @lc id=2463 lang=golang
// @title minimum-recolors-to-get-k-consecutive-black-blocks

package main

// @test("WBBWWBBWBW",7)=3
// @test("WBWBBBW",2)=0
// @test("WBBWWWWBBWWBBBBWWBBWWBBBWWBBBWWWBWBWW", 15)=6
// @test("WWBBBWBBBBBWWBWWWB", 16)=6
func minimumRecolors(blocks string, k int) int {
	pre := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			pre++
		}
	}
	result := pre
	for i := 0; i < len(blocks)-k; i++ {
		if blocks[i+k] == 'W' {
			pre++
		}
		if blocks[i] == 'W' {
			pre--
		}
		if pre < result {
			result = pre
		}
	}
	return result
}
