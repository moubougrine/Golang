package main

import "fmt"

func main() {
	res := ""
	result := ""
	for i := 'A'; i <= 'Z'; i++ {
		res += string(i)
	}
	for i, ele := range res {
		if (i+1)%2 == 0 {
			result += string(ele + 32)
		} else {
			result += string(ele)
		}
	}
	re := ""
	result1 := ""
	for j := 'Z'; j >= 'A'; j-- {
		re += string(j)
	}
	for j, ele1 := range re {
		if (j+1)%2 == 0 {
			result1 += (string(ele1 + 32))
		} else {
			result1 += (string(ele1))
		}
	}
	fmt.Println(result)
	fmt.Println(result1)
}
