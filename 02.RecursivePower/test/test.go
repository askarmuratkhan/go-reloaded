package main

import (
	"fmt"

	student ".."
)

func main() {
	arg1 := -7
	arg2 := -2
	fmt.Println(student.RecursivePower(arg1, arg2))
	arg3 := -8
	arg4 := -7
	fmt.Println(student.RecursivePower(arg3, arg4))
	arg5 := 4
	arg6 := 8
	fmt.Println(student.RecursivePower(arg5, arg6))
	arg7 := 1
	arg8 := 3
	fmt.Println(student.RecursivePower(arg7, arg8))
	arg9 := -1
	arg10 := 1
	fmt.Println(student.RecursivePower(arg9, arg10))
	arg11 := -6
	arg12 := 5
	fmt.Println(student.RecursivePower(arg11, arg12))
}
