package main

import "maps"

// Category: algorithms
// Level: Hard
// Percent: 34.82192%

// You are given a string s and an array of strings words. All the strings of words are of the same length.
//
// A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.
//
//
// 	For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
//
//
// Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.
//
//
// Example 1:
//
//
// Input: s = "barfoothefoobarman", words = ["foo","bar"]
//
// Output: [0,9]
//
// Explanation:
//
// The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
// The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
//
//
// Example 2:
//
//
// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
//
// Output: []
//
// Explanation:
//
// There is no concatenated substring.
//
//
// Example 3:
//
//
// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
//
// Output: [6,9,12]
//
// Explanation:
//
// The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
// The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
// The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].
//
//
//
// Constraints:
//
//
// 	1 <= s.length <= 10⁴
// 	1 <= words.length <= 5000
// 	1 <= words[i].length <= 30
// 	s and words[i] consist of lowercase English letters.
//

func findSubstring(s string, words []string) []int {
	m, n := len(words), len(words[0])
	if m == 0 || n > len(s) {
		return []int{}
	}

	result := make([]int, 0)

	dict := make(map[string]int)
	for _, w := range words {
		dict[w]++
	}

	for i := 0; i <= len(s)-m*n; i++ {
		match := true
		wordDict := maps.Clone(dict)

		for j := 0; j < m; j++ {
			candidate := s[i+j*n : i+j*n+n]
			cnt, ok := wordDict[candidate]
			if !ok {
				match = false
				break
			}
			if cnt == 1 {
				delete(wordDict, candidate)
			} else {
				wordDict[candidate]--
			}
		}

		if match {
			result = append(result, i)
		}
	}

	return result
}
