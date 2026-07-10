package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type TreeNode struct {
	Val      int
	Height   int
	Parents  []*TreeNode
	Children []*TreeNode
}

// given: N tasks and there dependency you must complete previous task one day before another
// 节点的高度由最长的路径决定，所以这个问题是看从根到叶子的最大高度
// 特殊情况，怎么判断有环： 从根节点出发，遍历的时候不会碰到重复的节点（简单粗暴的方法） 如果没有根节点怎么办？
// 怎么判断根节点： 没有父节点的
// 判断叶子节点： 没有子节点的
// 转为一个图问题
func getWorkingDayCount(N int, relations [][]int) int {
	// step1 build trees
	maxHeight := 0
	nodes := make([]*TreeNode, N)
	for i := 0; i < N; i++ {
		nodes[i] = &TreeNode{
			Val: i + 1,
		}
	}

	for j := 0; j < len(relations); j++ {
		// set parent and child
		p, c := relations[j][0], relations[j][1]
		if len(nodes[c-1].Parents) == 0 {
			nodes[c-1].Parents = []*TreeNode{
				nodes[p-1],
			}
		} else {
			nodes[c-1].Parents = append(nodes[c-1].Parents, nodes[p-1])
		}

		if len(nodes[p-1].Children) == 0 {
			nodes[p-1].Children = []*TreeNode{
				nodes[c-1],
			}
		} else {
			nodes[p-1].Children = append(nodes[p-1].Children, nodes[c-1])
		}
	}

	visited := 0
	// step2 loop get maxHeight of tree
	// filter root and dfs
	var dfs func(p *TreeNode, state int)
	dfs = func(p *TreeNode, state int) {
		if maxHeight < 0 || p == nil {
			return
		}

		visited = visited | 1<<p.Val - 1
		if len(p.Children) == 0 {
			return
		}

		for k := 0; k < len(p.Childrens); k++ {
			pk := p.Childrens[k]
			// check if already visited
			if 1<<(pk.Val-1)&state > 0 {
				maxHeight = -1
				return
			}

			// if not, try next state and update height
			pk.Height = max(p.Height+1, pk.Height)
			maxHeight = max(maxHeight, pk.Height)
			nextState := state | 1<<(pk.Val-1)
			dfs(pk, nextState)
			// restore state do nothing
		}

	}

	for i := 0; i < N; i++ {
		if len(nodes[i].Parents) == 0 {
			nodes[i].Val = 1
			dfs(nodes[i], 1<<i)
			if maxHeight < 0 {
				return -1
			}
		}
	}

	// step3 check all visited and return maxHeight
	if visisted+1 < 1<<N {
		return -1
	}

	return maxHeight
}
