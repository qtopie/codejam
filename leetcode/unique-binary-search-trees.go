package leetcode

// 当根节点确定时，左右子树的内容确定，分解为复合子问题
// recurrence:
//      f(j) = f(j-1) * f(i-j)
func numTrees(n int) int {
	g := make([]int, n+1)
	g[0] = 1
	g[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}

	return g[n]
}

func generateTrees(n int) []*TreeNode {
	g := make([][]*TreeNode, n+1)
	g[0] = []*TreeNode{nil}
	g[1] = []*TreeNode{&TreeNode{Val: 1}}

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// g[i] += g[j-1] * g[i-j]
			g[i] = append(g[i], generateTreeWithRoot(i, j, g)...)
		}
	}

	return g[n]
}

func uniqueTrees(n int) [][]string {
	results := [][]string{}
	trees := generateTrees(n)

	for _, t := range trees {
		results = append(results, t.ToArray())
	}
	return results
}

// set j as root, length is j
func generateTreeWithRoot(i, j int, g [][]*TreeNode) []*TreeNode {
	result := []*TreeNode{}

	if j == 1 {
		for m := 1; m <= len(g[i-1]); m++ {
			root := &TreeNode{Val: j}
			root.Right = copyTree(g[i-j][m-1], j)
			result = append(result, root)
		}
	} else if j == i {
		for l := 1; l <= len(g[j-1]); l++ {
			root := &TreeNode{Val: j}
			root.Left = g[j-1][l-1]
			result = append(result, root)
		}
	} else {
		// from 1 to j-1
		for l := 1; l <= len(g[j-1]); l++ {
			// from j+1 to i
			for m := 1; m <= len(g[i-j]); m++ {
				root := &TreeNode{Val: j}
				root.Left = g[j-1][l-1]
				root.Right = copyTree(g[i-j][m-1], j)
				result = append(result, root)
			}
		}
	}

	return result
}

func copyTree(root *TreeNode, adder int) *TreeNode {
	if root == nil {
		return nil
	}

	newRoot := &TreeNode{Val: root.Val + adder}
	newRoot.Left = copyTree(root.Left, adder)
	newRoot.Right = copyTree(root.Right, adder)
	return newRoot
}
