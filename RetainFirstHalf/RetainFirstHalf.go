package main

import "fmt"

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
func RetainFirstHalf(str string) string {
	result := ""
	for i,ele:= range str {
		result+=string(ele)
		if i == (len(str)-1)/2{
			break
		}
	}
	res:= ""
	for i,ele:= range result{
		if i ==len(result)-1 && ele == ' '{
			continue
		}else{
			res+=string(ele)
		}
	}
	return res
}
