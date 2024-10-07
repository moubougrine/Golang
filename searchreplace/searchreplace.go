package main

import (
	"fmt"
	"os"
)

func main() {
	are := os.Args[1]
	are1 := os.Args[2]
	are2:= os.Args[3]
	result := ""
	for _, ele :=range are{
		for _,ele1:=range are1{
			for _, ele2 := range are2{
				if ele == ele1{
					ele = ele2
					result+=string(ele)
				}else {
					result+=string(ele)
				}
			}
		}
	}
	fmt.Println(result)
}