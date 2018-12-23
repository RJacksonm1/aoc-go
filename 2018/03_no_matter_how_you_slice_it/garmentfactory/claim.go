package garmentfactory

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Claim to a piece of fabric
type Claim struct {
	ID   int
	X, Y int
	W, H int
}

// NewClaimFromDeclaration creates a Claim struct from a formatted claim string
func NewClaimFromDeclaration(declaration string) (c Claim, err error) {
	claimRegexp, err := regexp.Compile(`^#(\d+?) @ (\d+?),(\d+?): (\d+?)x(\d+?)$`)
	if err != nil {
		return Claim{}, err
	}

	matches := claimRegexp.FindStringSubmatch(declaration)

	if len(matches) == 0 {
		return Claim{}, errors.New("Claim declaration is incomplete or formatted incorrectly")
	}

	c.ID, _ = strconv.Atoi(matches[1])
	c.X, _ = strconv.Atoi(matches[2])
	c.Y, _ = strconv.Atoi(matches[3])
	c.W, _ = strconv.Atoi(matches[4])
	c.H, _ = strconv.Atoi(matches[5])

	return c, nil
}

// ToString represents the claim as a declaration string
func (c *Claim) ToString() string {
	return fmt.Sprintf("#%d @ %d,%d: %dx%d", c.ID, c.X, c.Y, c.W, c.H)
}
