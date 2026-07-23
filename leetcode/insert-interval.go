package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	low, high := newInterval[0], newInterval[1]

	if len(intervals) == 0 || low < intervals[0][0] && high > intervals[len(intervals)-1][1] {
		return [][]int{newInterval}
	} else if low > intervals[len(intervals)-1][1] {
		return append(intervals, newInterval)
	} else if high < intervals[0][0] {
		result = append(result, newInterval)
		result = append(result, intervals...)
		return result
	}

	// init index of intervals for newInterval
	l, r := -1, -1

	// try to find value in place
	l1, r1 := intervalsBinarySearch(intervals, low)
	if l1 == r1 {
		l = l1
	}

	l2, r2 := intervalsBinarySearch(intervals, high)
	if l2 == r2 {
		r = l2
	}

	// neither found in interval
	// [r1]l[l1], [r2]r[l2]
	if l == -1 && r == -1 {
		result = append(result, intervals[:l1]...)
		result = append(result, newInterval)
		result = append(result, intervals[l2:]...)
	} else if l == -1 {
		// left not found in intervals, while right is in one interval
		// [r1]l[l1], r2=r=l2
		intervals[r][0] = low
		result = append(result, intervals[:l1]...)
		result = append(result, intervals[r:]...)
	} else if r == -1 {
		// right not found in intervals, while left is in one interval
		// l1=r1=l, [r2]r[l2]
		intervals[l][1] = high
		result = append(result, intervals[:l+1]...)
		result = append(result, intervals[l2:]...)
	} else {
		// both found in intervals
		// l1=r1=l, l2=r2=r
		result = append(result, intervals[:l]...)
		merged := []int{intervals[l][0], intervals[r][1]}
		result = append(result, merged)
		result = append(result, intervals[r+1:]...)
	}

	return result
}

func intervalsBinarySearch(intervals [][]int, target int) (int, int) {
	l, r := 0, len(intervals)-1
	for l <= r {
		m := l + (r-l)/2
		if target >= intervals[m][0] && target <= intervals[m][1] {
			return m, m
		} else if target < intervals[m][0] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return l, r
}
