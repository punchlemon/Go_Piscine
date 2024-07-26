package main

import (
	"fmt"
	"piscine"
)

func main() {
	var strings []string
	var s string
	elems := []string{"Hello!", " How", " are", " you?"}
	sep := ":"
	fmt.Println(piscine.Join(elems, sep))
	fmt.Println(piscine.Join(strings, sep))
	strings = []string{"", "", "", s}
	fmt.Println(piscine.Join(strings, sep))
}
