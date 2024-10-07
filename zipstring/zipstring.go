package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func Tostring(s int) string {
	res := ""
	char := "0123456789"
	for s > 0 {
		res = string(char[s%10]) + res
		s /= 10
	}
	return res
}
func ZipString(s string) string {
	result := ""
	count := 1
	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1]{
			count++
		}else{
			result+=Tostring(count)+string(s[i-1])
			count=1
		}
	}
	return result
}
