package piscine

func ToUpper(s string) string {
	result := []rune(s)
	for i, c := range result {
		if c >= 'a' && c <= 'z' {
			result[i] = c - ('a' - 'A')
		}
	}
	return string(result)
}
