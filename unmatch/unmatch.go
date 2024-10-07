package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch :=Unmatch(a)
	fmt.Println(unmatch)
}
func Unmatch(a []int) int {
	count:=make(map[int]bool)
	result:=0
	for _,ele:=range a{
		if !count[ele]{
			result+=ele
			count[ele] = true
		}else{
			result-=ele
		}
	}
	return result
}