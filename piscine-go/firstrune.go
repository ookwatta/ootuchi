package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return 0 // Return 0 if the string is empty
}
