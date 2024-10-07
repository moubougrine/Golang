package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	are := os.Args[1]
	char, _ := strconv.Atoi(are)
	count := []string{}
	for i := 2; i <= char; i++ {
		for char > 1 && char%i == 0 {
			if IsPrime(char) {
				res := strconv.Itoa(i)
				count = append(count, res)
				char /= i
			}
		}
	}
	ress := ""
	for i, ele := range count {
		if i < len(count)-1 {
			ress += string(ele) + "*"
		} else {
			ress += string(ele)
		}
	}
	fmt.Println(ress)
}
func IsPrime(s int) bool {
	for i := 2; i <= s; i++ {
		if s%i == 0 {
			return true
		}
	}
	return false
}
