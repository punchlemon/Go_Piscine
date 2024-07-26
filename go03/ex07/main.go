package main

import (
	"fmt"
	"piscine"
)

func main() {
	var s string
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
	fmt.Println(piscine.Capitalize(""))
	fmt.Println(piscine.Capitalize(s))
}
