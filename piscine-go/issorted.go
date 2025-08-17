package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	// Check ascending order
	ascending := true
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			ascending = false
			break
		}
	}

	if ascending {
		return true
	}

	// Check descending order
	descending := true
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			descending = false
			break
		}
	}

	return descending
}
