package main

// Calibrate the chronal device with the list of possible corrections
// The correct frequency is the one we observe multiple times
func Calibrate(freqChanges []int) int {
	var frequency int
	var seenFreqs = make(map[int]bool, len(freqChanges))

	for {
		for _, correction := range freqChanges {
			if seenFreqs[frequency] {
				return frequency
			}
			seenFreqs[frequency] = true

			frequency += correction
		}
	}
}
