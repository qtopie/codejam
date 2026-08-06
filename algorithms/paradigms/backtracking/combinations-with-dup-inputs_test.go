package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func normalize(s [][]int) [][]int {
	out := make([][]int, len(s))
	for i := range s {
		c := append([]int(nil), s[i]...)
		sort.Ints(c)
		out[i] = c
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		la, lb := len(a), len(b)
		for k := 0; k < la && k < lb; k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return la < lb
	})
	return out
}

func equalSlices(a, b [][]int) bool {
	return reflect.DeepEqual(normalize(a), normalize(b))
}

func Test_combinationsWithDup(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		input []int
		want  [][]int
	}{
		{name: "empty", input: []int{}, want: [][]int{{}}},
		{name: "single", input: []int{1}, want: [][]int{{}, {1}}},
		{name: "duplicates-pair", input: []int{2, 2}, want: [][]int{{}, {2}, {2, 2}}},
		{name: "1-2-2", input: []int{1, 2, 2}, want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
		{name: "unsorted-2-1-2", input: []int{2, 1, 2}, want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationsWithDup(tt.input)
			if !equalSlices(got, tt.want) {
				t.Errorf("combinationsWithDup(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
