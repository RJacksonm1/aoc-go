package main

// Calibrate the chronal device with the
func Calibrate(freqChanges []int) int {
	var frequency int

	for _, correction := range freqChanges {
		frequency += correction
	}

	return frequency
}
