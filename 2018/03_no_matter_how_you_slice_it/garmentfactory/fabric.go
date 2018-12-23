package garmentfactory

import "fmt"

// Fabric is a 2d map (y,y) of square inches, where each inch could be
// claimed (many times) for use in a particular garment
type Fabric map[int]map[int][]Claim

// Conflict represents a square-inch of fabric with multiple claims
type Conflict struct {
	x, y   int
	claims []Claim
}

// AddClaim Add a claim to the fabric
func (f *Fabric) AddClaim(c Claim) {
	for x := c.X; x < (c.X + c.W); x++ {
		for y := c.Y; y < (c.Y + c.H); y++ {
			if (*f)[x] == nil {
				(*f)[x] = make(map[int][]Claim)
			}

			(*f)[x][y] = append((*f)[x][y], c)
		}
	}
}

// Conflicts will return the square inches of fabric which have multiple claims
func (f *Fabric) Conflicts() (c []Conflict) {
	for x, ys := range *f {
		for y, claims := range ys {
			if len(claims) <= 1 {
				continue
			}

			c = append(c, Conflict{x, y, claims})
		}
	}

	return c
}

// Width of the fabric, computed from the claims
func (f *Fabric) Width() int {
	biggestW := 0

	for _, ys := range *f {
		for _, claims := range ys {
			for _, c := range claims {
				if c.X+c.W > biggestW {
					biggestW = c.X + c.W
				}
			}
		}
	}

	return biggestW
}

// Height of the fabric, computed from the claims
func (f *Fabric) Height() int {
	biggestH := 0

	for _, ys := range *f {
		for _, claims := range ys {
			for _, c := range claims {
				if c.Y+c.H > biggestH {
					biggestH = c.Y + c.H
				}
			}
		}
	}

	return biggestH
}

// Render an illustration of the map to CLI
func (f *Fabric) Render() {
	w := f.Width()
	h := f.Height()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			claims := (*f)[x][y]
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
