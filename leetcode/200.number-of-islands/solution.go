package main

// Category: algorithms
// Level: Medium
// Percent: 64.286156%

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
//
//
// Example 1:
//
// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1
//
//
// Example 2:
//
// Input: grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// Output: 3
//
//
//
// Constraints:
//
//
// 	m == grid.length
// 	n == grid[i].length
// 	1 <= m, n <= 300
// 	grid[i][j] is '0' or '1'.
//

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {
				continue
			}
			// found a new island
			count++
			// iterative DFS using index = r*n + c
			stack := []int{i*n + j}
			grid[i][j] = '0'
			for len(stack) > 0 {
				idx := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				r := idx / n
				c := idx % n

				// neighbors: up, down, left, right
				if r-1 >= 0 && grid[r-1][c] == '1' {
					grid[r-1][c] = '0'
					stack = append(stack, (r-1)*n+c)
				}
				if r+1 < m && grid[r+1][c] == '1' {
					grid[r+1][c] = '0'
					stack = append(stack, (r+1)*n+c)
				}
				if c-1 >= 0 && grid[r][c-1] == '1' {
					grid[r][c-1] = '0'
					stack = append(stack, r*n+(c-1))
				}
				if c+1 < n && grid[r][c+1] == '1' {
					grid[r][c+1] = '0'
					stack = append(stack, r*n+(c+1))
				}
			}
		}
	}

	return count
}
