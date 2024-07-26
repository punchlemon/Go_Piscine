package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("010203"))
	fmt.Println(piscine.IsNumeric("01,02,03"))
	var s string
	fmt.Println(piscine.IsNumeric(""))
	fmt.Println(piscine.IsNumeric(s))
}
