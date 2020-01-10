package student

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {

	if n < 1 || n > 9 {
		Print("\n")
		return
	}

	current := getFirstNumber(n)

	lastNumber := getLastNumber(n)

	for current != lastNumber {
		Print(current + ", ")
		current = getNextNumber(current)
	}

	Print(current + "\n")
}

func getFirstNumber(n int) string {
	r := ""

	for i := 0; i < n; i++ {
		r += string(rune(i + '0'))
	}

	return r
}

func getLastNumber(n int) string {
	r := ""
	for i := 0; i < n; i++ {
		r = string(rune(9-i+'0')) + r
	}
	return r
}

func getNextNumber(curr string) string {

	current := []rune(curr)

	for {
		current = increment(current)
		if isSorted(current) {
			break
		}
	}

	return string(current)
}

func increment(s []rune) []rune {
	len := 0
	for range s {
		len++
	}

	if s[len-1] != '9' {
		s[len-1]++
	} else {
		f := increment(s[:len-1])
		return []rune(string(f) + "0")
	}

	return s
}

func isSorted(s []rune) bool {
	for i := range s {
		if i != 0 && s[i-1] >= s[i] {
			return false
		}
	}

	return true
}

func Print(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

/*
package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 1 && n > 9 {
		return
	}
	mx_ln := 1
	for i := 2; i <= n; i++ {
		mx_ln *= 10
	}
	for i := mx_ln / 10; i < mx_ln; i++ {
		if check(i) == true {
			if mx_ln >= 10 {
				z01.PrintRune('0')
			}
			GIVE(i)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	for i := mx_ln; i <= mx_ln*9; i++ {
		if check(i) == true {
			GIVE(i)
			if ok(i) == true {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func GIVE(y int) {
	w := '0'
	if y == 0 {
		z01.PrintRune(w)
		return
	}
	for j := 1; j <= y%10; j++ {
		w++
	}
	for j := -1; j >= y%10; j-- {
		w++
	}
	if y/10 != 0 {
		GIVE(y / 10)
	}
	z01.PrintRune(w)
	return
}

func check(p int) bool {
	for {
		if p == 0 {
			break
		}
		if p/10 != 0 && p%10 <= ((p/10)%10) {
			return false
		}
		p /= 10
	}
	return true
}
func ok(p int) bool {
	if p < 9 {
		return true
	}
	if p%10 == 9 {
		for {
			if p == 0 {
				break
			}
			if p/10 != 0 && p%10 != ((p/10)%10)+1 {
				return true
			}
			p /= 10
		}

		return false
	} else {
		return true
	}
}
*/
