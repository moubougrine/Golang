package main 

import "fmt"

func main(){
    fmt.Println(CountNegative([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
    fmt.Println(CountNegative([]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}))
    fmt.Println(CountNegative([]int{}))
    fmt.Println(CountNegative([]int{-1,2,0,-3}))
}
func CountNegative(numbers []int) int {
    count := 0 
	for _,ele := range numbers{
		if ele < 0{
			count++
		}
	}
	return count
}