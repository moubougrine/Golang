package piscine

import (
	"fmt"
)
func PrintCombN(n int) {
	if n == 1 {
		for i := 0; i <= 9; i++ {
			if i == 9 {
				fmt.Print(i)
			} else {
				fmt.Print(i, ",", " ")
			}
		}
	}
	if n == 3 {
		for i := 0; i <= 7; i++ {
			for j := i + 1; j <= 8; j++ {
				for k := j + 1; k <= 9; k++ {
					if !(i == 7 && j == 8 && k == 9) {
						fmt.Print(i)
						fmt.Print(j)
						fmt.Print(k)
						fmt.Print(",")
						fmt.Print(" ")
					} else {
						fmt.Print(i)
						fmt.Print(j)
						fmt.Print(k)
					}
				}
			}
		}
	}
	if n == 9 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				fmt.Print(j)
			}
			if i != 9 {
				fmt.Print(","," ")
			}
		}
	}
	fmt.Println()
}
