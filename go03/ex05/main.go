package main

import (
	"fmt"
	"piscine"
)

func main() {
	var s string
	fmt.Println(piscine.ToUpper("Hello! How are you?"))
	fmt.Println(piscine.ToUpper(""))
	fmt.Println(piscine.ToUpper(s))
}
