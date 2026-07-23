package leetcode

// 输入：s = "eceba", k = 2
// 输出：3
// 解释：满足题目要求的子串是 "ece" ，长度为 3
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) <= k {
		return len(s)
	}

	max := 0
	chs := []rune(s)

	// init window start pos and map
	l := 0
	dict := make(map[rune]int)

	for i := 0; i < len(chs); i++ {
		count, _ := dict[chs[i]]
		dict[chs[i]] = count + 1

		if len(dict) > k {
			dict[chs[l]] = dict[chs[l]] - 1
			l++

			if dict[chs[l]] == 0 {
				delete(dict, chs[l])
			}
		} else if (i - l + 1) > max {
			max = i - l + 1
		}
	}

	return max
}

// 读懂题目, (举例, 推导)
// 数学建模, 最优解法
// 伪代码 coding
