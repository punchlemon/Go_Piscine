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
		PrintNV()
		return len, true
	}
	for i, c := range base {
		for j := i+1; j < len; j++ {
			if (c == []rune(base)[j]) || (c == '-' || c == '+') {
				PrintNV()
				return len, true
			}
		}
	}
	return len, false
}

func PrintNbrBase(nbr int, base string) {
	base_len, err := CheckBase(base)
	if err {
		return
	}
	if nbr < 0 {
		ft.PrintRune('-')
	}
	for nbr != 0 {
		i := nbr%base_len
		if i < 0 {
			i *= -1
		}
		defer ft.PrintRune([]rune(base)[i])
		nbr /= base_len
	}
}
