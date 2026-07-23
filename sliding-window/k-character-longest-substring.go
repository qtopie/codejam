package sliding_window

// 题目：至多包含 K 个不同字符的最长子串
// 题目描述：
// 给定一个字符串 s 和一个整数 k，请你找出其中至多包含 k 个不同字符的最长子串，并返回该子串的长度。
// 示例：
// 输入: s = "eceba", k = 2
// 输出: 3
// 解释: 满足条件的子串是 "ece"，长度为 3
func longestKSubstring(s string, k int) int {
	if len(s) <= 0 {
		return 0
	}
	if k == 0 {
		return 0
	}

	chs := []rune(s)
	m := make(map[rune]int, 0)
	maxLen := 0
	// 窗口定义为 [l, r]里的所有元素, 初始化的时候从r开始扩展, r=0
	for l, r := 0, 0; r < len(s); r++ {
		m[chs[r]]++

		for len(m) > k {
			m[chs[l]]--
			if m[chs[l]] == 0 {
				delete(m, chs[l])
			}
			l++
		}

		maxLen = max(maxLen, r-l+1)
	}

	return maxLen
}
