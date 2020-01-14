package student

func AtoiBase(s string, base string) int {
	if !checkBase(base) {
		return 0
	}

	if !checkNumberForBase(s, base) {
		return 0
	}

	answer := 0
	lent := 0
	for range base {
		lent++
	}

	for _, char := range s {
		for index2, char2 := range base {
			if char == char2 {
				answer = answer*lent + index2
			}
		}
	}

	return answer
}

func checkNumberForBase(nbr, base string) bool {
	for _, n := range nbr {
		flag := false
		for _, b := range base {
			if b == n {
				flag = true
				break
			}
		}

		if !flag {
			return false
		}
	}

	return true
}

func checkBase(base string) bool {
	if StringLen(base) < 2 {
		return false
	}
	return Unique(base)
}

func Unique(s string) bool {
	runes := []rune(s)
	for i, ch := range runes {
		for j := i + 1; j < StringLen(string(runes)); j++ {
			if ch == runes[j] && i != j {
				return false
			}
		}
	}
	return Signs(s)
}

func Signs(s string) bool {

	for _, ch := range s {
		if ch == '+' || ch == '-' {
			return false
		}
	}

	return true
}
func StringLen(str string) int {
	i := 0
	for range str {
		i++
	}
	return i

}
