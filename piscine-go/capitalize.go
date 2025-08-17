package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if newWord && r >= 'a' && r <= 'z' {
				runes[i] = r - 32 // to uppercase
			} else if !newWord && r >= 'A' && r <= 'Z' {
				runes[i] = r + 32 // to lowercase
			}
			newWord = false
		} else {
			newWord = true
		}
	}
	return string(runes)
}
