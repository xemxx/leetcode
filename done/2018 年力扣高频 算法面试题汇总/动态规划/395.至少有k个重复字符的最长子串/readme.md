采用分治法，通过递归实现

核心思想：如果某个字符 x 在整个字符串中出现的次数 < k，那么 x 不可能出现在最终要求的子串中。因此，可以将原字符串截断为：
> x 左侧字符子串 + x + x 右侧字符子串

代码：
``` golang
func longestSubstring(s string, k int) int {
	l := len(s)
	if l == 0 || l < k {
		return 0
	}
	if k < 2 {
		return l
	}
	return find(s, k)
}

func find(s string, k int) int {
    l, r := 0, len(s)-1
    //如果当前总长度小于k，肯定不存在，直接返回
	if r-l+1 < k {
		return 0
    }
    //首先将所有字符的出现次数统计
    m := map[byte]int{}
	for _, k := range s {
		m[byte(k)]++
    }
    
    //如果在边缘的字符出现次数小于k，那么自然可以剔除此字符，等于直接求单独x另一边的字符串]
    //判断l>r的原因是防止越界
	for m[s[l]] < k && l>r {
		l++
	}
	for m[s[r]] < k && l<r {
		r--
    }
    //在剔除边缘字符后再次判断总长度
	if r-l+1 < k {
		return 0
    }
    //开始遍历剩下的字符串中字符的情况，即，如果i字符出现次数小于k，则求i的左子串和右子串中的最大值
	for i := l; i <= r; i++ {
		if m[s[i]] < k {
			left := find(s[l:i], k)
			right := find(s[i+1:r+1], k)
			if left < right {
				left = right
			}
			return left
		}
    }
    //如果在上面不存在i所在字符的出现次数小于k，则剩下的字符串满足所有字符大于出现次数k，直接返回即可
	return r - l + 1
}
```