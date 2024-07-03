package piscine

func NRune(s string, n int) rune {
	length := 0
	if s != "" {
		for range s {
			length++
		}
	}
	if n < 0 || n > length {
		return rune(0)
	}
	return []rune(s)[n-1]
}