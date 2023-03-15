// @algorithm @lc id=1756 lang=golang
// @title minimum-deletions-to-make-string-balanced

package main

// @test("aababbab")=2
// @test("bbaaaaabb")=2
func minimumDeletions(s string) int {
	// 贪心
	return minimumDeletions2(s)
	aOp := 0
	bCount := 0
	res := 0
	for _, v := range s {
		if v == 'b' {
			bCount++
		} else if bCount > 0 {
			aOp++
			if aOp >= bCount {
				res += aOp
				bCount = 0
				aOp = 0
			}
		}
	}
	return res + aOp
}

func minimumDeletions1(s string) int {
	cntA := 0
	for _, v := range s {
		if v == 'a' {
			cntA++
		}
	}
	cntB := 0
	result := cntA
	for _, v := range s {
		if v == 'a' {
			cntA--
		} else if v == 'b' {
			cntB++
		}
		result = min(result, cntB+cntA)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeletions2(s string) int {
	// 动态规划,最后一个是b,则等于上一个的最小，最后一个是a,则要么去掉a，即上一个最小加一，要么保留a,去掉前面的b，取最小值
	cntB := 0
	result := 0
	for _, v := range s {
		if v == 'b' {
			cntB++
		} else {
			result = min(result+1, cntB)
		}

	}
	return result
}
