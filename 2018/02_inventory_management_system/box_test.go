package main

import (
	"reflect"
	"testing"
)

var boxTestCases = []struct {
	input    Box
	expected map[rune]int
}{
	{
		input: Box{id: "abcdef"},
		expected: map[rune]int{
			'a': 1,
			'b': 1,
			'c': 1,
			'd': 1,
			'e': 1,
			'f': 1,
		},
	},
	{
		input: Box{id: "bababc"},
		expected: map[rune]int{
			'a': 2,
			'b': 3,
			'c': 1,
		},
	},
	{
		input: Box{id: "abbcde"},
		expected: map[rune]int{
			'a': 1,
			'b': 2,
			'c': 1,
			'd': 1,
			'e': 1,
		},
	},
	{
		input: Box{id: "abcccd"},

		expected: map[rune]int{
			'a': 1,
			'b': 1,
			'c': 3,
			'd': 1,
		},
	},
	{
		input: Box{id: "aabcdd"},
		expected: map[rune]int{
			'a': 2,
			'b': 1,
			'c': 1,
			'd': 2,
		},
	},
	{
		input: Box{id: "abcdee"},
		expected: map[rune]int{
			'a': 1,
			'b': 1,
			'c': 1,
			'd': 1,
			'e': 2,
		},
	},
	{
		input: Box{id: "ababab"},
		expected: map[rune]int{
			'a': 3,
			'b': 3,
		},
	},
}

func TestDuplicateLetterCounts(t *testing.T) {
	for _, testCase := range boxTestCases {
		if res := testCase.input.DuplicateLetterCounts(); !reflect.DeepEqual(res, testCase.expected) {
			t.Fatalf("FAIL: Input:%s\nExpected: %d\nActual: %d",
				testCase.input, testCase.expected, res)
		}
		t.Logf("PASS: %s", testCase.input)
	}
}

func BenchmarkDuplicateLetterCounts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range boxTestCases {
			tc.input.DuplicateLetterCounts()
		}
	}
}
