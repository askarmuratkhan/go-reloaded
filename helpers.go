package student

import "github.com/01-edu/z01"

// Print print string
func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

// PrintNbr print number
func PrintNbr(n int) {

	r := ""

	if n == 0 {
		z01.PrintRune('0')
	} else if n < 0 {
		Print("-")
		for n != 0 {
			r = string(rune((n%10*-1)+48)) + r
			n /= 10
		}
	} else if n > 0 {
		for n != 0 {
			r = string(rune((n%10)+48)) + r
			n /= 10
		}
	}

	Print(r)
}

func Atoi(s string) int64 {
	var result int64
	var letter rune
	var sign int64 = 1

	for index := range s {
		letter = rune(s[index])
		if letter == '-' && index == 0 {
			sign *= -1
		} else if letter == '+' && index == 0 {
			continue
		} else if letter < '0' || letter > '9' {
			return 0
		} else {
			var n int64 = 0
			for i := '0'; i < letter; i++ {
				n++
			}
			result = result*10 + n
		}
	}

	return result * sign
}
