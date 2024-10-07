package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	char := "0123456789"
	res := ""
	if n == 0 {
		res = "0"
	}
	if n < 0 {
		n *= -1
		for n > 0 {
			res = string(char[n%10]) + res
			n /= 10
		}
		res = "-" + res
	} else {
		for n > 0 {
			res = string(char[n%10]) + res
			n /= 10
		}
	}
	for _, ele := range res {
		z01.PrintRune(ele)
	}

}
