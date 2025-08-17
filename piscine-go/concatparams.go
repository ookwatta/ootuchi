package piscine

func ConcatParams(args []string) string {
	result := ""
	for i := 0; i < len(args); i++ {
		if i > 0 {
			result += "\n" // Add newline before arguments after the first one
		}
		result += args[i]
	}
	return result
}
