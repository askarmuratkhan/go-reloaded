package main

import (
	"fmt"

	student ".."
)

func main() {
	result1 := student.ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	result2 := student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	result3 := student.ConvertBase("256850", "0123456789", "01")
	result4 := student.ConvertBase("uuhoumo", "choumi", "Zone01")
	result5 := student.ConvertBase("683241", "0123456789", "0123456789")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
}
