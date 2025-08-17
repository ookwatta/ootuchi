package piscine

func NRune(s string, n int) rune {
	runes := []rune(s) // Convert string to slice of runes

	if n <= 0 || n > len(runes) {
		return 0 // Return 0 if n is out of bounds
	}

	return runes[n-1] // Return the n-th rune (1-based index)
}
