package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	j := 0
	for _, v := range s {
		if v != "" {
			s[j] = v
			j++
		}
	}
	*ptr = s[:j]
	return j
}
