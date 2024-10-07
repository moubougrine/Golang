package main

import (
	"fmt"
)

const N = 6

func Compact(ptr *[]string) int {
	p:= 0
	for i := 0 ; i <= N ; i++{
		if i%2!=0{
			continue
		}else{
			p=N/i
		}
	}
	return p
}

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:",Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}