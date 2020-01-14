package main

import (
	"fmt"

	student ".."
)

func main() {
	nbits1 := student.ActiveBits(15)
	nbits2 := student.ActiveBits(17)
	nbits3 := student.ActiveBits(4)
	nbits4 := student.ActiveBits(11)
	nbits5 := student.ActiveBits(9)
	nbits6 := student.ActiveBits(12)
	nbits7 := student.ActiveBits(2)

	fmt.Println(nbits1)
	fmt.Println(nbits2)
	fmt.Println(nbits3)
	fmt.Println(nbits4)
	fmt.Println(nbits5)
	fmt.Println(nbits6)
	fmt.Println(nbits7)
}
