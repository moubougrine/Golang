package main

import "fmt"

func main(){
    fmt.Println(ReverseStrings([]string{"a", "b", "c"}))
    fmt.Println(ReverseStrings([]string{"Good","Morning!"}))
    fmt.Println(ReverseStrings([]string{"Hello World"}))
}
func ReverseStrings(strs []string) string {
	result := ""
	res:=""
	for i,ele:= range strs{
		if i < len(strs)-1{
			result+=string(ele)+" "
		}else{
			result+=string(ele)
		}
	}
	for _,el:=range result{
		res=string(el)+res
	}
	return res
}