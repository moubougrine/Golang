package main

import (
	"fmt"
	"os"
)

func main() {
	are := os.Args[1]
	chek, p := isVowel(are)
	result := ""
	if chek {
		result += are[p:] + are[:p] + "ay"
	} else {
		result = "No vowels"
	}
	fmt.Println(result)
}
func isVowel(str string) (bool, int) {
	for i, ele := range str {
		if ele == 'a' || ele == 'e' || ele == 'i' || ele == 'o' || ele == 'u' || ele == 'A' || ele == 'E' || ele == 'I' || ele == 'O' || ele == 'U' {
			return true, i
		}
	}
	return false, 0
}
