package piscine

import "ft"

func PrintStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

func PrintNV() {
	PrintStr("NV")
}

func CheckBase(base string) (int, bool) {
	len := 0
	for range base {
		len++
	}
	if len < 2 {
		return len, true
	}
	for i, c := range base {
		for j := i+1; j < len; j++ {
			if (c == []rune(base)[j]) || (c == '-' || c == '+') {
				return len, true
			}
		}
	}
	return len, false
}

func AtoiBase(s string, base string) int {
	res := 0
	base_len, err := CheckBase(base)
	if err {
		return res
	}
	s_len := 0
	for range s {
		s_len++
	}
	for _, c := range s {
		for j, b := range base {
			if c == b {
				res *= base_len
				res += j
				break
			}
		}
	}
	return res
}
