package main

import (
	"fmt"
	"piscine"
)

func main() {
	var strings []string
	var s string
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(elems))
	fmt.Println(piscine.BasicJoin(strings))
	strings = []string{"", "a", "", s}
	fmt.Println(piscine.BasicJoin(strings))
}
