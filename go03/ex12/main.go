package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))
	var s string
	fmt.Println(piscine.IsPrintable(""))
	fmt.Println(piscine.IsPrintable(s))
}
