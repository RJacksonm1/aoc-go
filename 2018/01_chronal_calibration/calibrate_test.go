package main

import (
	"testing"
)

var testCases = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, -2, 3, 1},
		expected: 3,
	},
	{
		input:    []int{1, 1, 1},
		expected: 3,
	},
	{
		input:    []int{1, 1, -2},
		expected: 0,
	},
	{
		input:    []int{-1, -2, -3},
		expected: -6,
	},
}

func TestCalibrate(t *testing.T) {
	for _, testCase := range testCases {
		if res := Calibrate(testCase.input); res != testCase.expected {
			t.Fatalf("FAIL: Input:%d\nExpected: %d\nActual: %d",
				testCase.input, testCase.expected, res)
		}
		t.Logf("PASS: %d", testCase.input)
	}
}

func BenchmarkCalibrate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Calibrate(tc.input)
		}
	}
}
