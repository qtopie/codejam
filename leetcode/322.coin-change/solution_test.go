package main

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{1, 2, 5}, 0, 0},
		{[]int{186, 419, 83, 408}, 6249, 20},
		{[]int{2}, 3, -1},
		{[]int{1}, 2, 2},
		{[]int{3, 7, 11}, 20, 4},
	}
	for _, tt := range tests {
		if got := coinChange(tt.coins, tt.amount); got != tt.want {
			t.Errorf("coinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
		}
		if got := coinChangeSkip(tt.coins, tt.amount); got != tt.want {
			t.Errorf("coinChangeSkip(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
		}
	}
}

func FuzzCoinChange(f *testing.F) {
	f.Add([]byte{1, 2, 5}, byte(11))
	f.Add([]byte{2}, byte(3))
	f.Add([]byte{3, 7}, byte(20))

	f.Fuzz(func(t *testing.T, coins []byte, amount byte) {
		if len(coins) == 0 || amount > 100 {
			t.Skip()
		}
		intCoins := make([]int, len(coins))
		for i, c := range coins {
			if c == 0 {
				t.Skip()
			}
			intCoins[i] = int(c)
		}
		a := int(amount)
		r1 := coinChange(intCoins, a)
		r2 := coinChangeSkip(intCoins, a)
		if r1 != r2 {
			t.Fatalf("classic=%d skip=%d, coins=%v, amount=%d", r1, r2, intCoins, a)
		}
	})
}
