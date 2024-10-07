package main

import (
	"fmt"
	"os"
)
func main() {
	are := os.Args[1]
	are1 := os.Args[2]
	are2 := os.Args[3]
	if overFlow(toInt(are)) || overFlow(toInt(are2)) {
		return
	}
	if !Operator(are1) {
		return
	}
	if isnumber(are) || isnumber(are2) {
		return
	}
	resul := 0
	if are1 == "+" {
		resul += toInt(are) + toInt(are2)
	} else if are1 == "-" {
		resul += toInt(are) - toInt(are2)
	} else if are1 == "*" {
		resul += toInt(are) * toInt(are2)
	}
	if are1 == "/" {
		if toInt(are2) != 0 {
			resul += toInt(are) / toInt(are2)
		} else {
			fmt.Println("No division by 0")
			return
		}
	}
	if are1 == "%" {
		if toInt(are2) != 0 {

			resul += toInt(are) % toInt(are2)
		} else {
			fmt.Println("No modulo by 0")
			return
		}
	}
	fmt.Print(resul)
}
func overFlow(n int) bool {
	if n >= 9223372036854775807 || -9223372036854775808 > n {
		return true
	}
	return false
}
func toInt(str string) int {
	char := 0
	are := 0
	r := 1
	for i, ele := range str {
		if ele == '-' && i == 0 {
			r = -1
			continue
		}
		are = int(ele - '0')
		char = 10*char + are
		are = 0
	}
	return char * r
}
func Operator(str string) bool {
	if str == "+" || str == "-" || str == "/" || str == "*" || str == "%" {
		return true
	}
	return false
}
func isnumber(s string) bool {
	for _, ele := range s {
		if ele >= 'a' && ele <= 'z' || ele>='A'&&ele<='Z'{
			return true
		}
	}
	return false
}
