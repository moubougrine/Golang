package main

import "fmt"

func main() {
	res:=""
	for i := 'z'; i >='a'; i-- {
		res+=string(i)
	}
	ree:=""
	for i,ele:=range res{
		if i%2!=0 && i!=0{
			ree+=string(ele-32)
		}else{
			ree+=string(ele)
		}
	}
	fmt.Print(ree)
}