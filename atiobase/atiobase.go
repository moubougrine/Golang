package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	n := 0
	for _,ele := range s {
		i := 0
		for j,ele1:= range base{
			if ele1=='+'||ele1=='-'{
				return 0
			}
			for k,ele2:= range base {
				if ele2==ele1&& j!=k{
					return 0
				}
			}
			if ele==ele1{
				break
			}
			i++
		}
		n=n*len(base)+i
	}
	return n

}