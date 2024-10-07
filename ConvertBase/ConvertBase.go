package main

import (
	"fmt"
)

func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
func ConvertBase(nbr, baseFrom, baseTo string) string {
	r := 0
	for _, ele := range nbr {
		i := 0
		for _, ele1 := range baseFrom {
			if ele == ele1 {
				break
			}
			i++
		}
		r = r*len(baseFrom) + i
	}
	res:= ""
	for r >0{
		res=string(baseTo[r%len(baseTo)]) + res
		r/=len(baseTo)
	}
	return res 
}
