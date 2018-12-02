package main

import (
	"testing"
)

var checksumTestCases = []struct {
	input    []Box
	expected int
}{
	{
		input: []Box{
			Box{id: "abcdef"},
			Box{id: "bababc"},
			Box{id: "abbcde"},
			Box{id: "abcccd"},
			Box{id: "aabcdd"},
			Box{id: "abcdee"},
			Box{id: "ababab"},
		},
		expected: 12,
	},
}

func TestChecksum(t *testing.T) {
	for _, testCase := range checksumTestCases {
		if res := Checksum(testCase.input); res != testCase.expected {
			t.Fatalf("FAIL: Input:%s\nExpected: %d\nActual: %d",
				testCase.input, testCase.expected, res)
		}
		t.Logf("PASS: %s", testCase.input)
	}
}

func BenchmarkChecksum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range checksumTestCases {
			Checksum(tc.input)
		}
	}
}
