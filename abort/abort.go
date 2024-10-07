package main

import (
	"fmt"
)

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
func Abort(a, b, c, d, e int) int {
	res:=[]int{a,b,c,d,e}
	result := (a+b+c+d+e)/len(res)
	return result
}