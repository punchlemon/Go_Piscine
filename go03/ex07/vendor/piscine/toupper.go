package piscine

func Capitalize(s string) string {
	valid := true
	var result string
	for _, c := range s {
		if valid && (c >= 'a' && c <= 'z') {
			result += string(c+'A'-'a')
		} else {
			result += string(c)
		}
		if (c>='a' && c<='z') || (c>='A' && c<='Z') || (c>='0' && c<='9') {
			valid = false
		} else {
			valid = true
		}
	}
	return string(result)
}