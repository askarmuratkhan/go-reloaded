package main

import (
	"os"

	student ".."
)

func main() {
	args := os.Args
	if len(args) > 1 {

		BigString := ""
		for i, val := range args[1:] {
			if i+1 == len(args[1:]) {
				BigString += val + ""
			} else {
				BigString += val + " "
			}
		}
		VowelString := ""
		for _, value := range BigString {
			if value == 'a' || value == 'e' || value == 'i' || value == 'o' || value == 'u' ||
				value == 'A' || value == 'E' || value == 'I' || value == 'O' || value == 'U' {
				VowelString += string(value)
			}

		}
		LenVowelString := 0
		for range VowelString {
			LenVowelString++
		}

		BigStringVowelIndex := make([]int, LenVowelString)
		j := 0
		for i, value := range BigString {
			if value == 'a' || value == 'e' || value == 'i' || value == 'o' || value == 'u' ||
				value == 'A' || value == 'E' || value == 'I' || value == 'O' || value == 'U' {
				BigStringVowelIndex[j] = i
				j++

			}
		}
		BigStringRune := []rune(BigString)

		for index := range VowelString {
			BigStringRune[BigStringVowelIndex[LenVowelString-1-index]] = rune(VowelString[index])
		}
		student.Print(string(BigStringRune))
	}

	student.Print("\n")
}
