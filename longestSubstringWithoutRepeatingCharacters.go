package leetcodeByGo

// 无重复字符的最长字串

// 计算最大值
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	// 构建字符与其在字符串中出现位置的映射关系
	m := make(map[byte]int, len(s))
	left, right, res := 0, 0, 0

	for right < len(s) {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = right
		} else {
			if m[s[right]]+1 > left {
				left = m[s[right]] + 1
			}

			m[s[right]] = right
		}

		res = max(res, right-left+1)
		right++
	}

	return res
}
