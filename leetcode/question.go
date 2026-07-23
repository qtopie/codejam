package leetcode

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	l := 0
	r := len(people) - 1
	count := 0
	place := sort.SearchInts(people, limit)

	if place == 0 || 2 * people[l] > limit {
		return len(people)
	}

	if place <= r {
		r = place - 1
		count += len(people) - place
	}

	for l < r {
		if people[r]+people[l] <= limit {
			count++
			l++
			r--
		} else {
			r--
			count++
		}
	}

	if l == r {
		count++
	}

	return count
}
