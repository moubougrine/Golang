package main

import "fmt"

func main() {
	fmt.Print(convert(-43))
}

/* func Convert(s string)int {
	result := 0
	se := 1
	for i,char:= range s{
		if i == 0 && char == '-'{
			se = -1
			continue
		}
		result = result*10 + int(char-'0')
	}
	return result * se
} */
func convert(s int) string {
	res := ""
	for i := 0 ;i<= s ;i++{
		if i == (s/10){
			res+= string(i)
		}
		if i == (s%10){
			res+=string(i)
		}
	}
	return res
}
