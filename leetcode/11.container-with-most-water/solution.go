package main

// 双指针, 面积等于边界最小高度乘以长度，复合子问题(但非最优)
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	result := 0

	// init with two pointers
	for l, r := 0, len(height)-1; l < r; {
		area := 0

		// 利用单调性移动双指针, 左右两边都单调
		// 减一算法, 贪心算法 全局最优
		if height[l] > height[r] {
			area = (r - l) * height[r]
			r--
		} else {
			area = (r - l) * height[l]
			l++
		}

		result = max(result, area)
	}

	return result
}
