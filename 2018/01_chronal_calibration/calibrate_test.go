package main

import (
	"testing"
)

var testCases = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, -1},
		expected: 0,
	},
	{
		input:    []int{3, 3, 4, -2, -4},
		expected: 10,
	},
	{
		input:    []int{-6, 3, 8, 5, -6},
		expected: 5,
	},
	{
		input:    []int{7, 7, -2, -7, -4},
		expected: 14,
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
