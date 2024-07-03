package piscine

func LastRune(s string) rune {
	length := 0
	if s != "" {
		for range s {
			length++
		}
	}
	if length == 0 {
		return rune(0)
	}
	return []rune(s)[length-1]
}