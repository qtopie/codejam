package leetcode

import (
	"sort"
	"strconv"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	dict := make(map[int][][]int)
	// build map
	for _, c := range candidates {
		for i := 1; i*c <= target; i++ {
			sum := i * c

			cmb := make([]int, i)
			for j := 0; j < i; j++ {
				cmb[j] = c
			}

			list, ok := dict[sum]
			if !ok {
				list = make([][]int, 1)
				list[0] = cmb
				dict[sum] = list
			} else {
				dict[sum] = append(list, cmb)
			}

			for k, v := range dict {
				if k+sum > target {
					continue
				}

				newList := make([][]int, 0)
				for _, l := range v {
					newL := make([]int, 0)
					newL = append(newL, l...)
					newL = append(newL, cmb...)
					newList = append(newList, newL)
				}

				sum2 := k + sum
				list, ok = dict[sum2]
				if !ok {
					dict[sum2] = newList
				} else {
					dict[sum2] = append(list, newList...)
				}
			}

		}
	}

	results, ok := dict[target]
	if !ok {
		return [][]int{}
	}

	return _distinct(results)
}

func _distinct(results [][]int) [][]int {
	newResult := make([][]int, 0)
	dict := make(map[string]bool)

	for _, result := range results {
		sort.Ints(result)
		k := ""
		for _, n := range result {
			k = k + strconv.Itoa(n)
		}

		if _, ok := dict[k]; !ok {
			newResult = append(newResult, result)
		}
		dict[k] = true
	}

	return newResult
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	results := make([][]int, 0)
	states := make([]int, 0)

	var backtrack func(states, candidates []int, target int, start int)
	backtrack = func(states, candidates []int, target int, start int) {
		if target == 0 {
			result := make([]int, len(states))
			copy(result, states)
			results = append(results, result)
		} else if target < 0 {
			// pruning
			return
		}

		for i := start; i < len(candidates); i++ {
			// pruning duplicates
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			states = append(states, candidates[i])
			backtrack(states, candidates, target-candidates[i], i+1)
			states = states[:len(states)-1]
		}
	}

	backtrack(states, candidates, target, 0)
	return results
}
