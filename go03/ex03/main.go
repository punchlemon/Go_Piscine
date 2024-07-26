package main

import (
	"fmt"
	"piscine"
)

func main() {
	var a string
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
	fmt.Println(piscine.Index("Hello!", ""))
	fmt.Println(piscine.Index("Hello!", a))
}
