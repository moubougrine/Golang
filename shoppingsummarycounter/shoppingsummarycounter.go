package main

import (
	"fmt"
)

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
func ShoppingSummaryCounter(str string) map[string]int {
	ma := make(map[string]int)
	res := ""
	for _, ele := range str {
		if ele == ' ' {
			if res != "" {
				ma[res]++
				res = ""
			}
		} else {
			res += string(ele)
		}
	}
	if res != "" {
		ma[res]++
	}
	return ma
}
