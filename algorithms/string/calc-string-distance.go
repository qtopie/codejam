package stringalgo

// CalcStringDistance computes edit distance between s1[start1:end1+1] and s2[start2:end2+1] recursively.
func CalcStringDistance(s1, s2 string, start1, end1, start2, end2 int) int {
	if start1 > end1 {
		if start2 > end2 {
			return 0
		}
		return end2 - start2 + 1
	}

	if start2 > end2 {
		if start1 > end1 {
			return 0
		}
		return end1 - start1 + 1
	}

	if s1[start1] == s2[start2] {
		return CalcStringDistance(s1, s2, start1+1, end1, start2+1, end2)
	}

	t1 := CalcStringDistance(s1, s2, start1+1, end1, start2+1, end2) // replace
	t2 := CalcStringDistance(s1, s2, start1+1, end1, start2, end2)   // delete from s1
	t3 := CalcStringDistance(s1, s2, start1, end1, start2+1, end2)   // insert into s1

	minVal := t1
	if t2 < minVal {
		minVal = t2
	}
	if t3 < minVal {
		minVal = t3
	}

	return minVal + 1
}
