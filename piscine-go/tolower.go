package piscine

func ToLower(s string) string {
	result := []rune(s)
	for i, c := range result {
		if c >= 'A' && c <= 'Z' {
			result[i] = c + ('a' - 'A')
		}
	}
	return string(result)
}
