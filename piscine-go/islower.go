package piscine

// IsLower checks if the string contains only lowercase alphabetic characters.
func IsLower(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
