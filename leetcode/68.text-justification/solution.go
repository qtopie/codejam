package main

// Category: algorithms
// Level: Hard
// Percent: 51.771225%

// Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.
//
// You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
//
// Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.
//
// For the last line of text, it should be left-justified, and no extra space is inserted between words.
//
// Note:
//
//	A word is defined as a character sequence consisting of non-space characters only.
//	Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
//	The input array words contains at least one word.
//
// Example 1:
//
// Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
// Output:
// [
//    "This    is    an",
//    "example  of text",
//    "justification.  "
// ]
//
// Example 2:
//
// Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
// Output:
// [
//   "What   must   be",
//   "acknowledgment  ",
//   "shall be        "
// ]
// Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
// Note that the second line is also left-justified because it contains only one word.
//
// Example 3:
//
// Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
// Output:
// [
//   "Science  is  what we",
//
//	"understand      well",
//
//   "enough to explain to",
//   "a  computer.  Art is",
//   "everything  else  we",
//   "do                  "
// ]
//
// Constraints:
//
//	1 <= words.length <= 300
//	1 <= words[i].length <= 20
//	words[i] consists of only English letters and symbols.
//	1 <= maxWidth <= 100
//	words[i].length <= maxWidth
func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)

	i := 0
	for i < len(words) {
		// break by line
		j := i   // record last word of line
		cnt := 0 // count of words and basic spaces

		for {
			if j > i {
				cnt++
			}
			cnt += len(words[j])

			// exceed the limit
			if j < len(words)-1 && cnt+1+len(words[j+1]) > maxWidth {
				break
			}

			// break when no more elements
			if j >= len(words)-1 {
				break
			}
			j++
		}

		// handle last line
		if j == len(words)-1 {
			s := words[i]
			for k := i + 1; k <= j; k++ {
				s += " " + words[k]
			}

			// 修正 1：用 for 循环补齐末尾所有缺失的空格
			for len(s) < maxWidth {
				s += " "
			}
			result = append(result, s)
		} else if j == i {
			s := words[i]
			for l := len(s); l < maxWidth; l++ {
				s += " "
			}
			result = append(result, s)
		} else {
			leftover := maxWidth - cnt
			m, n := leftover/(j-i), leftover%(j-i)

			s := words[i]
			for k := i + 1; k <= j; k++ {
				// 修正 2：初始包含 1 个基础空格
				space := " "
				for l := 0; l < m; l++ {
					space += " "
				}
				if k-i <= n {
					space += " "
				}
				s += space + words[k]
			}
			result = append(result, s)
		}

		i = j + 1
	}

	return result
}
