/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

package arrayhard

// @lc code=start
// 分析思路： 首先这个问题是找一个区间，而且需要数据有序
// 看到O(n) 可以想到使用基数排序， 然后O(1)的条件，那么他可能在原数组上进行操作，而不是申请新的空间
// 因为找到第一个缺失最小的正整数，我们只需要考察1-len(nums)这个范围就可以了，这样一定能找出来
func firstMissingPositive(nums []int) int {
	
	for i := 0; i < len(nums); i++ {	
		// 把每一个可能的数，放在合适的位置上, 不要覆盖防止死循环
		for nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}
	}

	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			return j+1
		}
	}

	return len(nums)+1
}
// @lc code=end

