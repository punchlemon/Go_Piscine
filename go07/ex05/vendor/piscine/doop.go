package piscine

import (
	"fmt"
	"os"
)

func Doop() {
	args := os.Args
	if args.len != 4 {
		return
	}
	args_len := 0
	for range args {
		args_len++
	}
	if len(args) != 4 {
		return
	}
	a := args[1]
	op := args[2]
	b := args[3]
	if !isNumeric(a) || !isNumeric(b) {
		fmt.Println(0)
		return
	}
	if op == "+" {
		fmt.Println(Atoi(a) + Atoi(b))
	} else if op == "-" {
		fmt.Println(Atoi(a) - Atoi(b))
	} else if op == "*" {
		fmt.Println(Atoi(a) * Atoi(b))
	} else if op == "/" {
		if Atoi(b) == 0 {
			fmt.Println("No division by 0")
		} else {
			fmt.Println(Atoi(a) / Atoi(b))
		}
	} else if op == "%" {
		if Atoi(b) == 0 {
			fmt.Println("No modulo by 0")
		} else {
			fmt.Println(Atoi(a) % Atoi(b))
		}
	} else {
		fmt.Println(0)
	}
}
