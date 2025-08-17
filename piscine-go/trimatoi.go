package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundDigit := false

	for _, ch := range s {
		if ch == '-' && !foundDigit {
			sign = -1
		} else if ch >= '0' && ch <= '9' {
			foundDigit = true
			result = result*10 + int(ch-'0')
		}
	}

	return sign * result
}
