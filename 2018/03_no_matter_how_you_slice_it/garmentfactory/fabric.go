package garmentfactory

import (
	"fmt"
)

// Coordinate represents a square-inch of a given Fabric
type Coordinate struct {
	X, Y int
}

// Fabric is a 2d map (y,y) of square inches, where each inch could be
// claimed (many times) for use in a particular garment
type Fabric struct {
	Claims map[int]Claim

	// Computed properties
	coordinateIndex map[Coordinate][]Claim
}

// Conflict represents a square-inch of fabric with multiple claims
type Conflict struct {
	x, y   int
	claims []Claim
}

// AddClaim Add a claim to the fabric
func (f *Fabric) AddClaim(c Claim) {
	if f.Claims == nil {
		f.Claims = make(map[int]Claim)
	}
	f.Claims[c.ID] = c
	f.indexClaim(c)
}

func (f *Fabric) indexClaim(c Claim) {
	if f.coordinateIndex == nil {
		f.coordinateIndex = make(map[Coordinate][]Claim)
	}

	for x := c.X; x < (c.X + c.W); x++ {
		for y := c.Y; y < (c.Y + c.H); y++ {
			xy := Coordinate{x, y}
			f.coordinateIndex[xy] = append(f.coordinateIndex[xy], c)
		}
	}
}

// Conflicts will return the square inches of fabric which have multiple claims
// Since we're going through all the claims, we might as well find the ones which don't conflict as well
func (f *Fabric) Conflicts() (conflicts []Conflict, remainder []Claim) {
	claimsAndConflicts := make(map[Claim]bool)

	for xy, claims := range f.coordinateIndex {

		siblingsConflict := len(claims) > 1

		if siblingsConflict {
			conflicts = append(conflicts, Conflict{xy.X, xy.Y, claims})
		}

		for _, c := range claims {
			// If this claim's already been tagged as conflicting, make sure
			// we never inadvertently mark it otherwise
			if claimsAndConflicts[c] != true {
				claimsAndConflicts[c] = siblingsConflict
			}
		}
	}

	unconflictingClaims := make([]Claim, 0)
	for c, conflicts := range claimsAndConflicts {
		if !conflicts {
			unconflictingClaims = append(unconflictingClaims, c)
		}
	}

	return conflicts, unconflictingClaims
}

// Width of the fabric, computed from the claims
func (f *Fabric) Width() int {
	biggestW := 0

	for _, c := range f.Claims {
		if c.X+c.W > biggestW {
			biggestW = c.X + c.W
		}
	}

	return biggestW
}

// Height of the fabric, computed from the claims
func (f *Fabric) Height() int {
	biggestH := 0

	for _, c := range f.Claims {
		if c.Y+c.H > biggestH {
			biggestH = c.Y + c.H
		}
	}

	return biggestH
}

// Render an illustration of the map to CLI
func (f *Fabric) Render() {
	w := f.Width()
	h := f.Height()
	plottedClaims := f.coordinateIndex

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			claims := plottedClaims[Coordinate{x, y}]
			claimCount := len(claims)

			if claimCount == 0 {
				fmt.Print(" ")
				continue
			}

			if claimCount >= 2 {
				fmt.Print("X")
				continue
			}

			// Multi-digit IDs throw spacing off, so lets substitute with dashes
			// fmt.Print(claims[0].ID)
			fmt.Print("-")
		}
		fmt.Print("\n")
	}
}
