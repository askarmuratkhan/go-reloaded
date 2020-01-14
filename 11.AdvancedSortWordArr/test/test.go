package main

import (
	"fmt"

	student ".."
)

func main() {

	result1 := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	result2 := []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	result3 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	result4 := []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	result5 := []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	result6 := []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}

	student.AdvancedSortWordArr(result1, student.Compare)
	student.AdvancedSortWordArr(result2, student.Compare)
	student.AdvancedSortWordArr(result3, student.Compare)
	student.AdvancedSortWordArr(result4, student.Compare)
	student.AdvancedSortWordArr(result5, student.Compare)
	student.AdvancedSortWordArr(result6, student.Compare1)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
}
