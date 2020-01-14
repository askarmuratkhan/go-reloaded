package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("bcbbbbaab", "abc"))
	fmt.Println(AtoiBase("0001", "01"))
	fmt.Println(AtoiBase("00", "01"))
	fmt.Println(AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println(AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(AtoiBase("thisinputshouldnotmatter", "abca"))
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
