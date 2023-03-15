/*
 * @lc app=leetcode.cn id=1096 lang=golang
 *
 * [1096] 花括号展开 II
 */

// @lc code=start
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

// @lc code=end