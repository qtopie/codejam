package backtracking

import (
	"reflect"
	"testing"
)

func TestLookupWordByPhoneNumber(t *testing.T) {
	tests := []struct {
		name   string
		number string
		want   []string
	}{
		{
			name:   "two digits 23",
			number: "23",
			want:   []string{"AD", "AE", "AF", "BD", "BE", "BF", "CD", "CE", "CF"},
		},
		{
			name:   "empty",
			number: "",
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LookupWordByPhoneNumber(tt.number)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupWordByPhoneNumber(%q) = %v, want %v", tt.number, got, tt.want)
			}
		})
	}
}
