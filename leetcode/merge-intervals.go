package leetcode

// [[1,3],[2,6],[8,10],[15,18]]
// https://leetcode.com/problems/merge-intervals/
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	insertionSortIntervals(intervals)

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		interval := result[len(result)-1]
		if interval[1] >= intervals[i][0] {
			interval[1] = max(intervals[i][1], interval[1])
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}


func insertionSortIntervals(intervals [][]int) [][]int {
	// for 1 to len-1 to insert
	for i := 1; i < len(intervals); i++ {
		for j := i; j > 0 && intervals[j][0] < intervals[j-1][0]; j-- {
			intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
		}
	}

	return intervals
}
