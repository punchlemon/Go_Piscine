package piscine

func ToUpper(s string) string {
	var result string
	for _, c := range s {
		if (c >= 'a' && c <= 'z') {
			result += string(c+'A'-'a')
		} else {
			result += string(c)
		}
	}
	return string(result)
}