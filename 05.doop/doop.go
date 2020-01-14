package main

import (
	"os"

	student ".."
	"github.com/01-edu/z01"
)

func main() {
	lens := 0
	for range os.Args {
		lens++
	}
	// 1. проверка длины os.Args
	if lens != 4 {
		return
	}
	var a, b int64
	var answer string
	str1 := os.Args[1]
	str2 := os.Args[3]
	// 2. проверка на цифровые аргументы
	if isNumeric(str1) == false || isNumeric(str2) == false {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	// 3. перевод аргументов с строки в числа
	a = student.Atoi(str1)
	b = student.Atoi(str2)
	operator := os.Args[2]
	// 4. проверка на значения переполнения, которые меняются свои знаки и значения при переводе со строки в число
	if str1[0] != '-' && a < 0 || str1[0] == '-' && a > 0 || str2[0] != '-' && b < 0 || str2[0] == '-' && b > 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	if operator == "+" {
		// проверка результата на значения переполнения
		if a > 0 && b > 0 && a+b < 0 || a < 0 && b < 0 && a+b > 0 {
			answer = Itoa1(0)
		}
		answer = Itoa1(a + b)
	} else if operator == "-" {
		// проверка результата на значения переполнения
		if a < 0 && b > 0 && a-b > 0 || a > 0 && b < 0 && a-b < 0 {
			answer = Itoa1(0)
		}
		answer = Itoa1(a - b)

	} else if operator == "*" {
		answer = mult(a, b)

	} else if operator == "/" {
		//проверка результата на значения переполнения
		if a < 0 && b < 0 && a/b < 0 {
			answer = Itoa1(0)
		}
		if b == 0 {
			answer = "No division by 0"
		} else {
			answer = Itoa1(a / b)
		}
	} else if operator == "%" {
		if b == 0 {
			answer = "No remainder of division by 0"
		} else {
			answer = Itoa1(a % b)
		}

	} else {
		// во всех остальных случаях оператора
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	for _, v := range answer {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Itoa1(n int64) string {
	var s string
	sign := true

	if n == -9223372036854775808 {
		s = "0"
		return s
	}
	if n < 0 {
		n = n * -1
		sign = false
	}

	for {
		mod := n % 10
		s = string(mod+'0') + s
		n = n / 10
		if n == 0 {
			break
		}
	}

	if sign == false {
		s = "-" + s
	}
	return s
}

func mult(a, b int64) string {
	res := a * b
	// проверка результата на значения переполнения
	if res/a != b {
		return Itoa1(0)
	}
	return Itoa1(a * b)
}

func isNumeric(str string) bool {

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
