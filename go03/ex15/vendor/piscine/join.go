package piscine

func Join(strs []string, sep string) string {
	len := 0
	for range strs {
		len++
	}
	var result string
	for i, s := range strs {
		result += s
		if i < len-1 {
			result += sep
		}
	}
	return result
}