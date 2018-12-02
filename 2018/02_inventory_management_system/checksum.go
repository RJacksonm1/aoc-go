package main

// Checksum our warehouse by duplicate letters in the box IDs:
// Checksum = boxes with 2 dupe letters * boxes with 3 dupe letters
func Checksum(boxes []Box) int {
	var twoDupeCount, threeDupeCount int

	for _, box := range boxes {
		var twoDupes bool
		var threeDupes bool

		for _, letterCount := range box.DuplicateLetterCounts() {
			switch letterCount {
			case 2:
				twoDupes = true

			case 3:
				threeDupes = true
			}
		}

		if twoDupes {
			twoDupeCount++
		}

		if threeDupes {
			threeDupeCount++
		}
	}

	return twoDupeCount * threeDupeCount
}
