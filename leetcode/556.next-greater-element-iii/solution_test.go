package main

import "testing"

func TestDebug(t *testing.T) {
	n := 2147483476
	got := nextGreaterElement(n)
	t.Logf("got = %#v", got)
	t.Log("Edit this test case first when debugging.")
}

func TestSamples(t *testing.T) {
	t.Run("sample_1", func(t *testing.T) {
		n := 12
		got := nextGreaterElement(n)
		t.Logf("got = %#v", got)
	})
	t.Run("sample_2", func(t *testing.T) {
		n := 21
		got := nextGreaterElement(n)
		t.Logf("got = %#v", got)
	})
}
