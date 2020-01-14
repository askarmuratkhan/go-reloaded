package student

func SplitWhiteSpaces(str string) []string {

	size := 1
	index := 0
	count := 0

	var prev rune
	var str1 = ""

	for _, v := range str {
		if (prev != ' ' && prev != '\t' && prev != '\n') || (v != ' ' && v != '\t' && v != '\n') {
			str1 += string(v)
		}
		prev = v
	}

	for i := 0; i < StrLength(str1); i++ {
		if str1[i] == ' ' || str1[i] == '\t' || str1[i] == '\n' {
			size++
		}
	}

	result := make([]string, size)

	for i := 0; i < StrLength(str1); i++ {
		if str1[i] == ' ' || str1[i] == '\t' || str1[i] == '\n' {

			result[count] = result[count] + str1[index:i]
			index = i + 1
			count++

		}

	}

	result[count] += str1[index:]

	return result

}

func StrLength(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
