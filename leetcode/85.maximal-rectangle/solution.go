package main

// Category: algorithms
// Level: Hard
// Percent: 58.237724%

// Given a rows x cols binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.
//
// Example 1:
//
// Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// Output: 6
// Explanation: The maximal rectangle is shown in the above picture.
//
// Example 2:
//
// Input: matrix = [["0"]]
// Output: 0
//
// Example 3:
//
// Input: matrix = [["1"]]
// Output: 1
//
// Constraints:
//
//	rows == matrix.length
//	cols == matrix[i].length
//	1 <= rows, cols <= 200
//	matrix[i][j] is '0' or '1'.

// DP 按照行来递推 以第 $i$ 行为底的最大矩形，其边界必然受限于第 $i$ 行的 0 分布，且受限于它头顶上那几行（即 $i-1$ 行及以上）的边界分布
// 高度 $h[j]$：如果 matrix[i][j] == '1'，则 $h[j] = h[j] + 1$。否则，$h[j] = 0$。左边界 $l[j]$：受限于当前行的连续 '1' 范围，也受限于上一行该位置的 $l[j]$。$l[j] = \max(l[j], \text{current\_left\_boundary})$。右边界 $r[j]$：同理，$r[j] = \min(r[j], \text{current\_right\_boundary})$。
// 也可以用单调栈 
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	n := len(matrix[0])
	h := make([]int, n)
	l := make([]int, n)
	r := make([]int, n)
	for i := range r {
		r[i] = n // 初始化右边界为最右侧
	}

	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		// 记录当前行连续1的位置
		curLeft, curRight := 0, n

		// 更新高度
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				h[j]++
			} else {
				h[j] = 0
			}
		}

		// 更新左边界
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if curLeft > l[j] {
					l[j] = curLeft
				}
			} else {
				l[j] = 0
				curLeft = j + 1
			}
		}

		// 更新右边界
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if curRight < r[j] {
					r[j] = curRight
				}
			} else {
				r[j] = n
				curRight = j
			}
		}

		// 计算面积
		for j := 0; j < n; j++ {
			area := h[j] * (r[j] - l[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
