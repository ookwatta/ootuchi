package piscine

func StrRev(s string) string {
	var revstr string

	for _, letter := range s {
		revstr = string(letter) + revstr
	}
	return string(revstr)
}
