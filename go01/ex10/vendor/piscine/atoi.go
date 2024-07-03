package piscine

func Atoi(s string) int {
	sign := 1
	num := 0
	for i, c := range s {
		if i == 0 && (c == '+' || c == '-') {
			if c == '+' {
				sign = 1
			} else if c == '-' {
				sign = -1
			}
		} else if c < '0' || c > '9' {
			return 0
		} else {
			num = num * 10 + int(c - '0')
		}
	}
	return sign * num
}
