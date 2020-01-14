package student

import (
	student ".."
)

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		student.Print("NV")
		return
	}

	if nbr == 0 {
		student.Print(string(base[0]))
		return
	}

	baseLen := 0
	for range base {
		baseLen++
	}

	if nbr < 0 {
		r := ""
		for nbr != 0 {
			r = string(base[nbr%baseLen*-1]) + r
			nbr /= baseLen
		}
		student.Print("-" + r)
		return
	}

	r := ""
	for nbr != 0 {
		r = string(base[nbr%baseLen]) + r
		nbr /= baseLen
	}
	student.Print(r)
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
