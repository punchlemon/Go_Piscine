package piscine

import "ft"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			ft.PrintRune(rune(i / 10 + '0'))
			ft.PrintRune(rune(i % 10 + '0'))
			ft.PrintRune(' ')
			ft.PrintRune(rune(j / 10 + '0'))
			ft.PrintRune(rune(j % 10 + '0'))
			if i != 98 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
}
