package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrime, a)
	fmt.Println(result)

	// prev := 2
	// new := 2
	// max := 0
	// for i := 2; i < 9223372036854775807; i++ {
	// 	if piscine.IsPrime(i) {
	// 		prev = new
	// 		new = i
	// 		tmp := new - prev - 1
	// 		if tmp > max {
	// 			max = tmp
	// 			fmt.Printf("len: %d, prime: %d -> %d\n", max, prev, new)
	// 		}
	// 	}
	// }

	// count := 0
	// for i := 2; i < 9223372036854775807; i++ {
	// 	if piscine.IsPrime(i) {
	// 		count++
	// 		fmt.Printf("%d, ", i)
	// 	}
	// 	if count == 5 {
	// 		count = 0
	// 		fmt.Println()
	// 	}
	// }

	// fmt.Println(piscine.IsPrime(2147483647))
	// fmt.Println(piscine.IsPrime(63386494859))
	// fmt.Println(piscine.IsPrime(968387269999))
	// fmt.Println(piscine.IsPrime(1424064219931))
	// fmt.Println(piscine.IsPrime(93645133675451))
	// fmt.Println(piscine.IsPrime(98891523235339))//14digit true
	// fmt.Println(piscine.IsPrime(877220334477629))//15digit true
	// fmt.Println(piscine.IsPrime(877220334477629))//15digit true
	// fmt.Println(piscine.IsPrime(20797611784096529))//17digit true
	// fmt.Println(piscine.IsPrime(61597993*96486979))//8digit*8digit false
	// fmt.Println(piscine.IsPrime(793123965019646797))//18digit
	// fmt.Println(piscine.IsPrime(2305843009213693951))
}
