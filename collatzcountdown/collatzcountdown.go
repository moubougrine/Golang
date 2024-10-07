package main

import (
	"fmt"
)

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}
func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	result := 0
	for start != 1 { 
		if start%2 == 0 {
			start /= 2
		} else {
			start = 3*start + 1
		}
		result++
	}
	return result
}
