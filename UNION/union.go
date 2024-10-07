package main

import (
	"os"
	"fmt"
)

func main() {
	are := os.Args[1]
	Are := os.Args[2]
	res := make(map[rune]bool)
	Union := func(str string){
		result:= ""
		for _, ele := range str {
			if !res[ele] {
				result+=string(ele)
				res[ele] = true
			}
		}
		fmt.Print(result)
	}
	Union(are)
	Union(Are)
}