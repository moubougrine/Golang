package main

import (
	"fmt"
)

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}

func DealAPackOfCards(deck []int) {
	for i := range deck{
		if i == 0 {
			fmt.Print("Player 1: ")
		}else if i == 3 {
			fmt.Print("Player 2: ")
		}else if i == 6 {
			fmt.Print("Player 3: ")
		}else if i == 9 {
			fmt.Print("Player 4: ")
		}
		if (i+1)% 3 == 0 {
			fmt.Println(deck[i])
		}else{
			if i < len(deck) {
				fmt.Print(deck[i], ","," ")
			}else{
				fmt.Print(deck[i])
			}
		}
	}
}