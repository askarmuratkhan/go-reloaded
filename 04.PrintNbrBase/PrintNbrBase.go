package student

import (
	"github.com/01-edu/z01"
)

// PrintNbrBase prints number nbr in base string
func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		Print("NV")
		return
	}

	// if nbr is zero
	if nbr == 0 {
		Print(string(base[0]))
		return
	}

	// calculate length of base
	baseLen := 0
	for range base {
		baseLen++
	}

	// if nbr is negative
	if nbr < 0 {
		r := ""
		for nbr != 0 {
			r = string(base[nbr%baseLen*-1]) + r
			nbr /= baseLen
		}
		Print("-" + r)
		return
	}

	// if nbr is positive
	r := ""
	for nbr != 0 {
		r = string(base[nbr%baseLen]) + r
		nbr /= baseLen
	}
	Print(r)
}

func isValidBase(s string) bool {
	strLen := 0
	sRune := []rune(s)
	for _, l := range sRune {
		if l == '+' || l == '-' {
			return false
		}
		strLen++
	}
	if strLen < 2 {
		return false
	}
	for i, l := range sRune {
		for j := i + 1; j < strLen; j++ {
			if l == sRune[j] {
				return false
			}
		}
	}
	return true
}

func Print(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// func ParseInt(n int) bool {
// 	if -9223372036854775808 <= n && n <= 9223372036854775807 {
// 		return true
// 	}
// 	return false
// }
