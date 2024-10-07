package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	res := ""
	for _, ss := range s {
		if IsUper(string(ss)){
			if IsUper(string(int(ss)+14)){
				res+=string(int(ss)+14)
			}else{
				res+=string(int(ss)-12)
			}
		}else if Islower(string(ss)){
				if Islower(string(int(ss)+14)){
				res+=string(int(ss)+14)
			}else{
				res+=string(int(ss)-12)
			}
		}else{
			res+=string(ss)
		}
	}
	return res
}
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
func Islower(s string)bool{
	for _,ss:= range s {
		if ss >= 'a' && ss<='z'{
			return true
		}
	}
	return false
}
func IsUper(s string)bool{
		for _,ss:= range s {
		if ss >= 'A' && ss<='Z'{
			return true
		}
	}
	return false
}