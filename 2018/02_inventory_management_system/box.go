package main

// A Box will have an id!
type Box struct {
	id string
}

// DuplicateLetterCounts is pretty self explanatory. It's used to compute a warehouse checksum
func (b *Box) DuplicateLetterCounts() map[rune]int {
	var letterCounts = make(map[rune]int, len(b.id))

	for _, l := range b.id {
		letterCounts[l]++
	}

	return letterCounts
}
