package main

import (
	"fmt"
	"os"
)

func main() {
	signs := []string{"+", "*", "-", "/", "%"}
	args := os.Args[1:]

	if LenArgs(args) != 3 {

	} else if !IsNumeric(args[0]) || !IsNumeric(args[2]) {
		fmt.Println("0")
	} else {
		if isSign(args[1], signs) {

			nb1 := atoi(args[0])
			nb2 := StrToInt(args[2])

			switch args[1] {
			case "+":
				fmt.Println(nb1 + nb2)
			case "-":
				fmt.Println(nb1 - nb2)
			case "/":
				if nb2 == 0 {
					fmt.Println("No division by 0")
				} else {
					fmt.Println(nb1 / nb2)
				}
			case "%":
				if nb2 == 0 {
					fmt.Println("No remainder of division by 0")
				} else {
					fmt.Println(nb1 % nb2)
				}
			case "*":
				fmt.Println(nb1 * nb2)

			}
		} else {
			fmt.Println("0")
		}
	}

}

func isSign(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func LenArgs(args []string) int {
	result := 0
	for range args {
		result++
	}
	return result
}

func StrToInt(s string) int {
	Rune := []rune(s)
	result := 0
	sign := 1

	for _, r := range Rune {
		count := 0
		if Rune[0] == '-' {
			sign = -1
		}
		for i := '0'; i < r; i++ {
			count++
		}
		result = 10*result + count

	}

	return result * sign
}

func IsNumeric(str string) bool {

	s := []rune(str)

	for i := 0; i < StrLen(str); i++ {
		if !IsNbr(s[i]) {
			return false
		}
	}
	return true

}

func IsNbr(s rune) bool {
	if s >= '0' && s <= '9' || s == '-' {
		return true
	}
	return false

}

func StrLen(str string) int {
	var result int
	ByteStr := []rune(str)

	for index := range ByteStr {
		result = index + 1
	}

	return result
}
