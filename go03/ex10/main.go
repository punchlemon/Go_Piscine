package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsUpper("HELLO"))
	fmt.Println(piscine.IsUpper("HELLO!"))
	var s string
	fmt.Println(piscine.IsUpper(""))
	fmt.Println(piscine.IsUpper(s))
}
