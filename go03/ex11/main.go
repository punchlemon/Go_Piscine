package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsLower("hello"))
	fmt.Println(piscine.IsLower("hello!"))
	var s string
	fmt.Println(piscine.IsLower(""))
	fmt.Println(piscine.IsLower(s))
}
