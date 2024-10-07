package main

import "fmt"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
}
func SortWordArr(a []string) {
	res := []string{}
	for _, ele := range a {
		if ele >= "0" && ele <= "9" {
			res = append(res, ele)
		}
	}
	for _, ele := range a {
		if ele >= "A" && ele <= "Z" {
			res = append(res, ele)
		}
	}
	for _, ele := range a {
		if ele >= "a" && ele <= "z" {
			res = append(res, ele)
		}
	}

	fmt.Print(res)
}
