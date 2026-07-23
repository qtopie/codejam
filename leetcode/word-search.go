package leetcode

func exist(board [][]byte, word string) bool {
	chs := make([]byte, 0)
	for _, r := range word {
		chs = append(chs, byte(r))
	}

	// for each line
	for i, row := range board {
		// find matched start
		for j, r := range row {
			if r == chs[0] {
				visited := make([]int, len(board))
				visited[i] = 1 << (len(row) - 1 - j)
				ok := dfsWordSearch(board, i, j, chs[1:], visited)
				if ok {
					return true
				}
			}
		}

	}

	return false
}

func dfsWordSearch(board [][]byte, i, j int, chs []byte, visited []int) bool {
	if len(chs) == 0 {
		return true
	}

	rowSize, columnSize := len(board), len(board[0])

	// up
	if i >= 1 {
		if board[i-1][j] == chs[0] && ((1<<columnSize|visited[i-1])&(1<<(columnSize-1-j))) == 0 {
			newVisited := make([]int, len(visited))
			copy(newVisited, visited)
			newVisited[i-1] = newVisited[i-1] | 1<<(columnSize-1-j)
			ok := dfsWordSearch(board, i-1, j, chs[1:], newVisited)
			if ok {
				return true
			}
		}
	}

	// down
	if i <= rowSize-2 {
		if board[i+1][j] == chs[0] && ((1<<columnSize|visited[i+1])&(1<<(columnSize-1-j))) == 0 {
			newVisited := make([]int, len(visited))
			copy(newVisited, visited)
			newVisited[i+1] = newVisited[i+1] | 1<<(columnSize-1-j)
			ok := dfsWordSearch(board, i+1, j, chs[1:], newVisited)
			if ok {
				return true
			}
		}
	}

	// bit calculation is wrong
	// left
	if j >= 1 {
		if board[i][j-1] == chs[0] && ((1<<columnSize|visited[i])&(1<<(columnSize-j)) == 0) {
			newVisited := make([]int, len(visited))
			copy(newVisited, visited)
			newVisited[i] = newVisited[i] | (1 << (columnSize - j))
			ok := dfsWordSearch(board, i, j-1, chs[1:], newVisited)
			if ok {
				return true
			}
		}
	}

	// right
	if j <= columnSize-2 {
		if board[i][j+1] == chs[0] && ((1<<columnSize|visited[i])&(1<<(columnSize-j-2)) == 0) {
			newVisited := make([]int, len(visited))
			copy(newVisited, visited)
			newVisited[i] = newVisited[i] | (1 << (columnSize - j - 2))
			ok := dfsWordSearch(board, i, j+1, chs[1:], newVisited)
			if ok {
				return true
			}
		}
	}

	return false
}
