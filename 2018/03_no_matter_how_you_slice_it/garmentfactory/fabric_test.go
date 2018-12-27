package garmentfactory

import (
	"bufio"
	"log"
	"os"
	"testing"
)

var testCases = []struct {
	fabric                 Fabric
	loadClaimsFromFile     string
	claims                 []Claim
	squareInchesInConflict int
	validClaimIds          map[int]bool
	setup                  bool
}{
	{
		fabric: Fabric{},
		claims: []Claim{
			MustCreateClaim(NewClaimFromDeclaration("#1 @ 1,3: 4x4")),
			MustCreateClaim(NewClaimFromDeclaration("#2 @ 3,1: 4x4")),
			MustCreateClaim(NewClaimFromDeclaration("#3 @ 5,5: 2x2")),
		},
		squareInchesInConflict: 4,
		validClaimIds: map[int]bool{
			3: true,
		},
	},
	{
		fabric:                 Fabric{},
		loadClaimsFromFile:     "../input.txt",
		claims:                 []Claim{},
		squareInchesInConflict: 106501,
		validClaimIds: map[int]bool{
			632: true,
		},
	},
}

// SetupTestCases loads any large test scenarios from files
func SetupTestCases() {
	for i := range testCases {
		if testCases[i].setup {
			continue
		}

		if testCases[i].loadClaimsFromFile == "" {
			for _, claim := range testCases[i].claims {
				testCases[i].fabric.AddClaim(claim)
			}
			testCases[i].setup = true
			continue
		}

		inputFile, err := os.Open(testCases[i].loadClaimsFromFile)
		if err != nil {
			log.Fatal(err)
		}
		defer inputFile.Close()

		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			claim, err := NewClaimFromDeclaration(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			testCases[i].fabric.AddClaim(claim)
			testCases[i].setup = true
		}
	}
}

func TestConflicts(t *testing.T) {
	SetupTestCases()

	for i, testCase := range testCases {
		// These exports suck... how to make consistent?
		conflicts, validClaims := testCase.fabric.Conflicts()
		squareInchesInConflict := len(conflicts)
		if squareInchesInConflict != testCase.squareInchesInConflict {
			t.Fatalf("FAIL: Test case %d failed conflict check\nExpected conflict: %d\nActual: %d",
				i, testCase.squareInchesInConflict, squareInchesInConflict)
		}

		for _, claim := range validClaims {
			if valid := testCase.validClaimIds[claim.ID]; !valid {
				t.Fatalf("FAIL: Test case %d failed valid claims check", i)
			}
		}
		t.Logf("PASS: %d", i)
	}
}

func BenchmarkConflicts(b *testing.B) {
	SetupTestCases()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			tc.fabric.Conflicts()
		}
	}
}
