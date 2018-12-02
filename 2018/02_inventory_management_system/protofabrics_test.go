package main

import (
	"testing"
)

var protoFabricsTestCases = []struct {
	input     []Box
	expectedA Box
	expectedB Box
}{
	{
		input: []Box{
			Box{id: "abcde"},
			Box{id: "fghij"},
			Box{id: "klmno"},
			Box{id: "pqrst"},
			Box{id: "fguij"},
			Box{id: "axcye"},
			Box{id: "wvxyz"},
		},
		expectedA: Box{id: "fghij"},
		expectedB: Box{id: "fguij"},
	},
}

func TestFindPrototypeFabrics(t *testing.T) {
	for _, testCase := range protoFabricsTestCases {
		if a, b := FindPrototypeFabrics(testCase.input); a.id != testCase.expectedA.id || b.id != testCase.expectedB.id {
			t.Fatalf("FAIL: Input:%s\nExpected: %s, %s\nActual: %s, %s",
				testCase.input, testCase.expectedA.id, testCase.expectedB.id, a.id, b.id)
		}
		t.Logf("PASS: %s", testCase.input)
	}
}

func BenchmarkFindPrototypeFabrics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range protoFabricsTestCases {
			FindPrototypeFabrics(tc.input)
		}
	}
}
