package main

import (
	"fmt"
	"piscine"
)

func main() {
	var s string
	fmt.Println(piscine.Concat("Hello!", " Hou are you?"))
	fmt.Println(piscine.Concat("", " Hou are you?"))
	fmt.Println(piscine.Concat(s, " Hou are you?"))
}
