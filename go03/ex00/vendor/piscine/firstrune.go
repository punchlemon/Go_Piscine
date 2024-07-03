package piscine

func FirstRune(s string) rune {
	if s == "" {
		return rune(0)
	}
	return []rune(s)[0]
}