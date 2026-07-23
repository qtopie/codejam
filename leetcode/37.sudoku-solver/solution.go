package main

import "math/bits"

// solveSudoku: 使用位掩码 + MRV（最少剩余值）启发式的高效数独求解器。
//
// 设计要点（实现思路）：
// - 用 9 个整型位掩码跟踪每一行、每一列、每个 3x3 宫中已被占用的数字：
//     rows[i], cols[j], boxes[b] 的第 d 位为 1 表示数字 d+1 已存在。
// - 将所有空格位置收集到 `empties` 列表，递归时我们按 MRV 选择下一个填充格：
//     在未填格中选择候选数字最少的格子以减少搜索分支（Minimum Remaining Value）。
// - 对某个格子可填的数字集合用位掩码表示（9 位，0x1FF 有效位），候选数的数量
//   可通过 `bits.OnesCount` 快速获得；枚举候选使用 `lsb := mask & -mask` 提取最低位。
// - 通过位运算快速放置/撤销数字：设置/清除对应行/列/宫的位，同时修改 `board`。

func solveSudoku(board [][]byte) {
	// rows/cols/boxes 的每一位代表数字 1..9（位 0 表示 '1'）
	var rows, cols, boxes [9]int
	empties := make([]int, 0, 81) // 存储空格位置，位置编码为 r*9+c

	// 初始化：填充掩码并记录空位
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				empties = append(empties, r*9+c)
			} else {
				d := int(board[r][c] - '1')
				bit := 1 << d
				rows[r] |= bit
				cols[c] |= bit
				boxes[(r/3)*3+c/3] |= bit
			}
		}
	}

	// dfs(idx) 负责填充 empties[idx:]
	var dfs func(int) bool
	dfs = func(idx int) bool {
		// 全部填完，终止条件
		if idx == len(empties) {
			return true
		}

		// MRV：从未处理的空位（idx..end）中挑选候选数最少的位置
		minIdx := -1
		minCount := 10
		candMask := 0
		for k := idx; k < len(empties); k++ {
			pos := empties[k]
			r := pos / 9
			c := pos % 9
			b := (r/3)*3 + c/3

			// 候选为行/列/宫都没有出现的数（取反并掩掉高位）
			mask := ^(rows[r] | cols[c] | boxes[b]) & 0x1FF
			cnt := bits.OnesCount(uint(mask))
			if cnt == 0 {
				// 无候选 -> 立即剪枝
				return false
			}
			if cnt < minCount {
				minCount = cnt
				minIdx = k
				candMask = mask
				if cnt == 1 {
					break // 最优（唯一候选），无需继续扫描
				}
			}
		}

		// 将选中的空位交换到当前 idx 位置（就地重排 empties），便于递归处理
		empties[idx], empties[minIdx] = empties[minIdx], empties[idx]
		pos := empties[idx]
		r := pos / 9
		c := pos % 9
		b := (r/3)*3 + c/3

		// 枚举候选位：使用低位提取（lsb）并通过 TrailingZeros 得到数字下标 d
		mask := candMask
		for mask != 0 {
			lsb := mask & -mask
			d := bits.TrailingZeros(uint(lsb)) // d in [0..8] 对应字符 '1'+d

			// 放置数字 d+1：设置三个位掩码并写回 board
			bit := 1 << d
			rows[r] |= bit
			cols[c] |= bit
			boxes[b] |= bit
			board[r][c] = byte('1' + d)

			// 递归填下一个空位
			if dfs(idx + 1) {
				return true
			}

			// 回溯（撤销）：清除位掩码并恢复为空格
			rows[r] &^= bit
			cols[c] &^= bit
			boxes[b] &^= bit
			board[r][c] = '.'

			// 从 mask 中移除已尝试的最低位，继续枚举下一个候选
			mask &= mask - 1
		}

		// 恢复之前的 empties 顺序（swap back），返回失败给上一层
		empties[idx], empties[minIdx] = empties[minIdx], empties[idx]
		return false
	}

	_ = dfs(0)
}
