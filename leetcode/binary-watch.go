package leetcode

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	if turnedOn == 0 {
		return []string{"0:00"}
	}

	hVals := _generateValues(4, "")
	hDigits := make([][]int, 5)
	hDigits[0] = []int{0}
	for _, s := range hVals {
		chs := []rune(s)
		sum, count := 0, 0
		for i := 0; i < len(chs); i++ {
			if chs[i] == '1' {
				sum = sum*2 + 1
				count++
			} else {
				sum = sum * 2
			}
		}

		if count > 0 && sum < 12 {
			hDigits[count] = append(hDigits[count], sum)
		}
	}

	mVals := _generateValues(6, "")
	mDigits := make([][]int, 7)
	mDigits[0] = []int{0}
	for _, s := range mVals {
		chs := []rune(s)
		sum, count := 0, 0
		for i := 0; i < len(chs); i++ {
			if chs[i] == '1' {
				sum = sum*2 + 1
				count++
			} else {
				sum = sum * 2
			}
		}

		if count > 0 && sum < 60 {
			mDigits[count] = append(mDigits[count], sum)
		}
	}

	result := make([]string, 0)
	for i := 0; i <= turnedOn && i <= 4; i++ {
		if turnedOn-i >= len(mDigits) {
			continue
		}

		left := hDigits[i]
		right := mDigits[turnedOn-i]
		for _, h := range left {
			for _, m := range right {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return result
}

func _generateValues(n int, s string) []string {
	if n == len(s) {
		return []string{s}
	}

	result := make([]string, 0)
	// n is 0
	vals0 := _generateValues(n, s+"0")
	// n is 1
	vals1 := _generateValues(n, s+"1")

	result = append(result, vals0...)
	result = append(result, vals1...)
	return result
}
