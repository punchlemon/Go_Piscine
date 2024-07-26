package piscine

import "ft"

func PrintNbr(n int) {
	if n < 0 {
		ft.PrintRune('-')
	}
	var c rune
	for ;n > 0; n /= 10 {
		if n < 0 {
			c = rune(-n%10 + '0')
		} else {
			c = rune(n%10 + '0')
		}
		defer ft.PrintRune(c)
	}
}