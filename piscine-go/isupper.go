package piscine

func IsUpper(s string) bool {
	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') {
			return false
		}
		if r >= 'a' && r <= 'z' {
			return false
		}
	}
	return true
}
