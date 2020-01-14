package student

func Atoi(s string) int {
	var result int
	var letter rune
	var sign int = 1

	for index := range s {
		letter = rune(s[index])
		if letter == '-' && index == 0 {
			sign *= -1
		} else if letter == '+' && index == 0 {
			continue
		} else if letter < '0' || letter > '9' {
			return 0
		} else {
			var n int = 0
			for i := '0'; i < letter; i++ {
				n++
			}
			result = result*10 + n
		}
	}

	return result * sign
}
