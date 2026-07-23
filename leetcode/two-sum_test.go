package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case 1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1}},
		},
		{"test2", args{[]int{2, 2, 2, 2, 2}, 8}, [][]int{
			{2, 2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourSumCount(t *testing.T) {
	// Define test cases in a table
	var testCases = []struct {
		name     string
		input1   []int
		input2   []int
		input3   []int
		input4   []int
		expected int
	}{
		{name: "Empty arrays", input1: []int{}, input2: []int{}, input3: []int{}, input4: []int{}, expected: 0},
		{name: "Example with solutions", input1: []int{1, 2}, input2: []int{-2, -1}, input3: []int{-1, 2}, input4: []int{0, 2}, expected: 2},
		{name: "All zeros", input1: []int{0, 0}, input2: []int{0, 0}, input3: []int{0, 0}, input4: []int{0, 0}, expected: 6}, // Combination of (0, 0, 0, 0)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fourSumCount(tc.input1, tc.input2, tc.input3, tc.input4)

			if actual != tc.expected {
				t.Errorf("Test case '%s': Expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}
