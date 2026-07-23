package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}

	resultCount := 0
	_dfs(&resultCount, make(map[int]int), root)
	return resultCount
}

func _dfs(resultCount *int, digitMap map[int]int, p *TreeNode) {
	if p == nil {
		return
	}

	digitMap[p.Val]++
	if p.Left == nil && p.Right == nil {
		if _checkValidSolution(digitMap) {
			*resultCount = *resultCount + 1
		}
		return
	}

	if p.Left != nil {
		_dfs(resultCount, digitMap, p.Left)
		digitMap[p.Left.Val]--
		if digitMap[p.Left.Val] == 0 {
			delete(digitMap, p.Left.Val)
		}
	}

	if p.Right != nil {
		_dfs(resultCount, digitMap, p.Right)
		digitMap[p.Right.Val]--
		if digitMap[p.Right.Val] == 0 {
			delete(digitMap, p.Right.Val)
		}
	}

}

func _checkValidSolution(dictMap map[int]int) bool {
	if len(dictMap) == 0 {
		return true
	}

	totalCount := 0
	oddCount := 0
	for _, count := range dictMap {
		if count%2 != 0 {
			oddCount++
		}
		totalCount += count
	}

	if totalCount%2 == 0 && oddCount > 0 {
		return false
	}

	if totalCount%2 == 1 && oddCount != 1 {
		return false
	}

	return true
}
