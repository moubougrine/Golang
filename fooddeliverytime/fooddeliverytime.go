package main

import (
	"fmt"
)

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	preptime:=0
	if order=="burger"{
		preptime=15
	}else if order == "chips"{
		preptime=10
	}else if order == "nuggets"{
		preptime = 12
	}
	return preptime
}

func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}
