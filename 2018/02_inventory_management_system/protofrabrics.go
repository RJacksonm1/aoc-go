package main

// FindPrototypeFabrics finds the two boxes whose IDs differ only by a single character
// Thankfully there's only 2, otherwise this would be much harder...
func FindPrototypeFabrics(boxes []Box) (x Box, y Box) {
	for _, a := range boxes {
		for _, b := range boxes {
			// Don't care about IDs with differing lengths
			if len(a.id) != len(b.id) {
				continue
			}

			var runesDifferent int
			for i := range a.id {
				if a.id[i] != b.id[i] {
					runesDifferent++
				}
			}

			if runesDifferent == 1 {
				return a, b
			}
		}
	}

	return x, y
}
