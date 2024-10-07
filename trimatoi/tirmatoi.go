package main

import (
	"fmt"
)

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
func TrimAtoi(s string) int {
	res := ""
	ress := ""
	resutl := 0
	for _, ele := range s {
		if ele > '0' && ele <= '9' || ele == '-' {
			res += string(ele)
		}
	}
	for i, els := range res {
		if els == '-' && i != 0 {
			continue
		} else {
			ress += string(els)
		}
	}
	r := 1
	for i, ell := range ress {
		if ell == '-' && i == 0 {
			r = -1
			continue
		}
		resutl = resutl*10 + int(ell-'0')
	}
	return resutl * r
}
