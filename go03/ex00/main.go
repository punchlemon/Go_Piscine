package main

import (
	"ft"
	"piscine"
)

func main() {
	var a string
	ft.PrintRune(piscine.FirstRune("Hello!"))
	ft.PrintRune(piscine.FirstRune("Salut!"))
	ft.PrintRune(piscine.FirstRune("Ola!"))
	ft.PrintRune(piscine.FirstRune(""))
	ft.PrintRune(piscine.FirstRune(a))
	ft.PrintRune('\n');
}
