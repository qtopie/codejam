package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, s := range strs {
		chs := []rune(s)
		sort.Slice(chs, func(i, j int) bool {
			return chs[i] > chs[j]
		})

		key := string(chs)
		if anagrams, ok := dict[key]; ok {
			dict[key] = append(anagrams, s)
		} else {
			dict[key] = []string{s}
		}
	}

	results := make([][]string, 0)
	for _, v := range dict {
		results = append(results, v)
	}
	return results
}
