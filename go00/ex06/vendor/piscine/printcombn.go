package piscine

import "ft"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	first := true
	GenerateCombinations("", 0, n, &first)
	ft.PrintRune('\n')
}

func GenerateCombinations(current string, start, n int, first *bool) {
	if len(current) == n {
		if !*first {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		} else {
			*first = false
		}
		for _, c := range current {
			ft.PrintRune(c)
		}
		return
	}
	for i := start; i < 10; i++ {
		next := current + string(rune(i + '0'))
		GenerateCombinations(next, i + 1, n, first)
	}
}
