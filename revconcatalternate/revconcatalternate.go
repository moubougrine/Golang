package main

import (
	"fmt"
)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
	fmt.Println(RevConcatAlternate([]int{}, []int{1, 2, 5, 3}))

}
func RevConcatAlternate(slice1, slice2 []int) []int {
	rs := []int{}
	if len(slice2)==0{
		rs = append(rs, slice1...)
	}
	if len(slice1)==0{
		rs = append(rs, slice2...)
	}
	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice2); j++ {
			if i == j {
				rs = append(rs, slice2[j], slice1[i])
			}
			if j >= len(slice1) {
				if j < len(rs)-1 {
					rs = append(rs, slice2[j])
				}
			}
			if i >= len(slice2) {
				rs = append(rs, slice1[i])
				break
			}
		}
	}
	return swap(rs)
}
func swap(r []int) []int {
	result := []int{}
	for i := 1; i <= len(r); i++ {
		result = append(result, r[len(r)-i])
	}
	return result
}
