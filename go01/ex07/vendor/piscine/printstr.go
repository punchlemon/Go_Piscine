package piscine

func StrRev(s string) string {
	runes := []rune(s)
	len := len(runes)
	for i := 0; i < len / 2; i++ {
		runes[i], runes[len - i - 1] = runes[len - i - 1], runes[i]
	}
	return string(runes)
}
